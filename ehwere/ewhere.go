package ewhere

import (
	"regexp"
	"strings"
)

func Parse(query string, params map[string]interface{}) (string, []interface{}) {
	re := regexp.MustCompile(`\?(\w+)`)
	matches := re.FindAllStringSubmatch(query, -1)

	args := []interface{}{}

	for _, match := range matches {
		fullPlaceholder := match[0]
		field := match[1]

		val, ok := params[field]
		if !ok || val == nil || val == "" {
			// Kalau kosong, replace ke '__PLACEHOLDER__' (yang nantinya jadi '1=1')
			query = strings.Replace(query, fullPlaceholder, "__PLACEHOLDER__", 1)
		} else {
			// Kalau ada, ganti jadi field = ?
			query = strings.Replace(query, fullPlaceholder, field+" = ?", 1)
			args = append(args, val)
		}
	}

	// Replace "__PLACEHOLDER__" jadi "1=1" setelah semua parsing
	query = strings.ReplaceAll(query, "__PLACEHOLDER__", "1=1")

	// Cleanup biasa
	query = strings.ReplaceAll(query, "WHERE 1=1 AND ", "WHERE ")
	query = strings.ReplaceAll(query, "WHERE 1=1 OR ", "WHERE ")
	query = strings.ReplaceAll(query, "AND 1=1", "")
	query = strings.ReplaceAll(query, "OR 1=1", "")
	query = strings.ReplaceAll(query, "(1=1)", "")

	return strings.TrimSpace(query), args
}
