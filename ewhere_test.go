package ewhere

import (
	"reflect"
	"strings"
	"testing"
)

func Test_ewhere(t *testing.T) {
	tests := []struct {
		name      string
		query     string
		params    map[string]interface{}
		wantQuery string
		wantArgs  []interface{}
	}{
		{
			name:  "Semua ada",
			query: "SELECT * FROM users WHERE ?name AND ?age",
			params: map[string]interface{}{
				"name": "Jane",
				"age":  25,
			},
			wantQuery: "SELECT * FROM users WHERE name = ? AND age = ?",
			wantArgs:  []interface{}{"Jane", 25},
		},
		{
			name:  "Partial kosong",
			query: "SELECT * FROM users WHERE ?name AND ?age",
			params: map[string]interface{}{
				"name": "",
				"age":  30,
			},
			wantQuery: "SELECT * FROM users WHERE age = ?",
			wantArgs:  []interface{}{30},
		},
		{
			name:      "Semua kosong",
			query:     "SELECT * FROM users WHERE ?name AND ?age",
			params:    map[string]interface{}{},
			wantQuery: "SELECT * FROM users WHERE 1=1",
			wantArgs:  []interface{}{},
		},
		{
			name:  "Bawaan tetap",
			query: "SELECT * FROM users WHERE name = 'Jane' AND ?age",
			params: map[string]interface{}{
				"age": 25,
			},
			wantQuery: "SELECT * FROM users WHERE name = 'Jane' AND age = ?",
			wantArgs:  []interface{}{25},
		},
		{
			name: "Multi-line query",
			query: `
SELECT id, name
FROM users
WHERE ?name
  AND (?age OR ?city)
`,
			params: map[string]interface{}{
				"name": "Jane",
				"age":  25,
				"city": "New York",
			},
			wantQuery: `
SELECT id, name
FROM users
WHERE name = ?
  AND (age = ? OR city = ?)
`,
			wantArgs: []interface{}{"Jane", 25, "New York"},
		},
		{
			name: "Nested SELECT",
			query: `
SELECT *
FROM (
    SELECT id
    FROM employees WHERE ?name
) AS sub
WHERE ?department
`,
			params: map[string]interface{}{
				"name":       "Jane",
				"department": "Model",
			},
			wantQuery: `
SELECT *
FROM (
    SELECT id
    FROM employees WHERE name = ?
) AS sub
WHERE department = ?
`,
			wantArgs: []interface{}{"Jane", "Model"},
		},
		{
			name:  "Field dengan titik",
			query: "SELECT * FROM products WHERE ?pr.code AND ?pr.category",
			params: map[string]interface{}{
				"pr.code":         "P001",
				"pr.category":     "Gadget",
				"testParamsLebih": "Banyak",
			},
			wantQuery: "SELECT * FROM products WHERE pr.code = ? AND pr.category = ?",
			wantArgs:  []interface{}{"P001", "Gadget"},
		},
		{
			name:  "Placeholder first in parentheses",
			query: "SELECT * FROM users WHERE (?name AND ?age)",
			params: map[string]interface{}{
				"name": "",
				"age":  30,
			},
			wantQuery: "SELECT * FROM users WHERE (age = ?)",
			wantArgs:  []interface{}{30},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotQuery, gotArgs := Parse(tt.query, tt.params)

			gotQueryNorm := normalize(gotQuery)
			wantQueryNorm := normalize(tt.wantQuery)

			if gotQueryNorm != wantQueryNorm {
				t.Errorf("got query = %v, want %v", gotQueryNorm, wantQueryNorm)
			}

			if !reflect.DeepEqual(gotArgs, tt.wantArgs) {
				t.Errorf("got args = %v, want %v", gotArgs, tt.wantArgs)
			}
		})
	}
}

func normalize(s string) string {
	return strings.TrimSpace(s)
}
