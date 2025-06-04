package ewhere

import (
	"reflect"
	"regexp"
	"strings"
)

var placeholderRE = regexp.MustCompile(`\?([\w\.]+)`)

// handleSlice builds an IN clause for a slice and appends the slice values
// to args. It returns the updated query string.
func handleSlice[T any](query, placeholder, field string, slice []T, args *[]any) string {
	if len(slice) == 0 {
		return strings.Replace(query, placeholder, "__PLACEHOLDER__", 1)
	}

	placeholders := strings.TrimSuffix(strings.Repeat("?,", len(slice)), ",")
	for _, v := range slice {
		*args = append(*args, v)
	}

	return strings.Replace(query, placeholder, field+" IN ("+placeholders+")", 1)
}

// Parse replaces dynamic placeholders in SQL with real fields and arguments.
//
// Example:
//
//	Input Query: "SELECT * FROM users WHERE ?name AND ?age"
//	Input Params: map[string]any{"name": "Jane", "age": 25}
//
//	Output Query: "SELECT * FROM users WHERE name = ? AND age = ?"
//	Output Args:  ["Jane", 25]
//
// Rules:
// - Placeholder format is `?field`.
// - If the param value is nil or empty (""), it will be ignored (replaced by "1=1").
// - Cleanup is applied to remove unnecessary "1=1" from the query.
//
// Special Notes:
// - Handles multi-line queries safely.
// - Cleans up leftover "AND 1=1", "OR 1=1", and "(1=1)".
//
// This function is designed to support dynamic SQL generation safely.
func Parse(query string, params map[string]any) (string, []any) {
	re := placeholderRE
	matches := re.FindAllStringSubmatch(query, -1)

	args := []any{}

	for _, match := range matches {
		fullPlaceholder := match[0]
		field := match[1]

		val, ok := params[field]
		if !ok || val == nil {
			// If missing or empty, temporarily replace with '__PLACEHOLDER__'
			query = strings.Replace(query, fullPlaceholder, "__PLACEHOLDER__", 1)
			continue
		}

		switch v := val.(type) {
		case string:
			if v == "" {
				query = strings.Replace(query, fullPlaceholder, "__PLACEHOLDER__", 1)
			} else {
				query = strings.Replace(query, fullPlaceholder, field+" = ?", 1)
				args = append(args, v)
			}
		default:
			rv := reflect.ValueOf(val)
			if rv.Kind() == reflect.Slice {
				slice := make([]any, rv.Len())
				for i := 0; i < rv.Len(); i++ {
					slice[i] = rv.Index(i).Interface()
				}
				query = handleSlice(query, fullPlaceholder, field, slice, &args)
			} else {
				query = strings.Replace(query, fullPlaceholder, field+" = ?", 1)
				args = append(args, val)
			}
		}
	}

	// After parsing, replace all '__PLACEHOLDER__' with '1=1'
	query = strings.ReplaceAll(query, "__PLACEHOLDER__", "1=1")

	// Cleanup unnecessary '1=1' fragments
	query = strings.ReplaceAll(query, "WHERE 1=1 AND ", "WHERE ")
	query = strings.ReplaceAll(query, "WHERE 1=1 OR ", "WHERE ")
	query = strings.ReplaceAll(query, "AND 1=1", "")
	query = strings.ReplaceAll(query, "OR 1=1", "")
	query = strings.ReplaceAll(query, "1=1 AND ", "")
	query = strings.ReplaceAll(query, "1=1 OR ", "")
	query = strings.ReplaceAll(query, "(1=1)", "")

	return strings.TrimSpace(query), args
}
