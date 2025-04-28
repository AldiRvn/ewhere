package ewhere

import (
	"regexp"
	"strings"
)

// Parse replaces dynamic placeholders in SQL with real fields and arguments.
//
// Example:
//   Input Query: "SELECT * FROM users WHERE ?name AND ?age"
//   Input Params: map[string]interface{}{"name": "Jane", "age": 25}
//
//   Output Query: "SELECT * FROM users WHERE name = ? AND age = ?"
//   Output Args:  ["Jane", 25]
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
func Parse(query string, params map[string]interface{}) (string, []interface{}) {
	re := regexp.MustCompile(`\?(\w+)`)
	matches := re.FindAllStringSubmatch(query, -1)

	args := []interface{}{}

	for _, match := range matches {
		fullPlaceholder := match[0]
		field := match[1]

		val, ok := params[field]
		if !ok || val == nil || val == "" {
			// If missing or empty, temporarily replace with '__PLACEHOLDER__'
			query = strings.Replace(query, fullPlaceholder, "__PLACEHOLDER__", 1)
		} else {
			// If present, replace with 'field = ?'
			query = strings.Replace(query, fullPlaceholder, field+" = ?", 1)
			args = append(args, val)
		}
	}

	// After parsing, replace all '__PLACEHOLDER__' with '1=1'
	query = strings.ReplaceAll(query, "__PLACEHOLDER__", "1=1")

	// Cleanup unnecessary '1=1' fragments
	query = strings.ReplaceAll(query, "WHERE 1=1 AND ", "WHERE ")
	query = strings.ReplaceAll(query, "WHERE 1=1 OR ", "WHERE ")
	query = strings.ReplaceAll(query, "AND 1=1", "")
	query = strings.ReplaceAll(query, "OR 1=1", "")
	query = strings.ReplaceAll(query, "(1=1)", "")

	return strings.TrimSpace(query), args
}
