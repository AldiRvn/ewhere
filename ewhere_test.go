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
		params    map[string]any
		wantQuery string
		wantArgs  []any
	}{
		{
			name:  "Semua ada",
			query: "SELECT * FROM users WHERE ?name AND ?age",
			params: map[string]any{
				"name": "Jane",
				"age":  25,
			},
			wantQuery: "SELECT * FROM users WHERE name = ? AND age = ?",
			wantArgs:  []any{"Jane", 25},
		},
		{
			name:  "Partial kosong",
			query: "SELECT * FROM users WHERE ?name AND ?age",
			params: map[string]any{
				"name": "",
				"age":  30,
			},
			wantQuery: "SELECT * FROM users WHERE age = ?",
			wantArgs:  []any{30},
		},
		{
			name:      "Semua kosong",
			query:     "SELECT * FROM users WHERE ?name AND ?age",
			params:    map[string]any{},
			wantQuery: "SELECT * FROM users WHERE 1=1",
			wantArgs:  []any{},
		},
		{
			name:  "Bawaan tetap",
			query: "SELECT * FROM users WHERE name = 'Jane' AND ?age",
			params: map[string]any{
				"age": 25,
			},
			wantQuery: "SELECT * FROM users WHERE name = 'Jane' AND age = ?",
			wantArgs:  []any{25},
		},
		{
			name: "Multi-line query",
			query: `
SELECT id, name
FROM users
WHERE ?name
  AND (?age OR ?city)
`,
			params: map[string]any{
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
			wantArgs: []any{"Jane", 25, "New York"},
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
			params: map[string]any{
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
			wantArgs: []any{"Jane", "Model"},
		},
		{
			name:  "Field dengan titik",
			query: "SELECT * FROM products WHERE ?pr.code AND ?pr.category",
			params: map[string]any{
				"pr.code":         "P001",
				"pr.category":     "Gadget",
				"testParamsLebih": "Banyak",
			},
			wantQuery: "SELECT * FROM products WHERE pr.code = ? AND pr.category = ?",
			wantArgs:  []any{"P001", "Gadget"},
		},
		{
			name:  "Placeholder first in parentheses",
			query: "SELECT * FROM users WHERE (?name AND ?age)",
			params: map[string]any{
				"name": "",
				"age":  30,
			},
			wantQuery: "SELECT * FROM users WHERE (age = ?)",
			wantArgs:  []any{30},
		},
		{
			name:  "Slice string",
			query: "SELECT * FROM users WHERE ?ids",
			params: map[string]any{
				"ids": []string{"A", "B", "C"},
			},
			wantQuery: "SELECT * FROM users WHERE ids IN (?,?,?)",
			wantArgs:  []any{"A", "B", "C"},
		},
		{
			name:  "Slice int",
			query: "SELECT * FROM users WHERE ?ids",
			params: map[string]any{
				"ids": []int{1, 2, 3},
			},
			wantQuery: "SELECT * FROM users WHERE ids IN (?,?,?)",
			wantArgs:  []any{1, 2, 3},
		},
		{
			name:  "Slice any",
			query: "SELECT * FROM users WHERE ?ids",
			params: map[string]any{
				"ids": []any{"A", "B", "C"},
			},
			wantQuery: "SELECT * FROM users WHERE ids IN (?,?,?)",
			wantArgs:  []any{"A", "B", "C"},
		},
		{
			name:  "Slice string empty",
			query: "SELECT * FROM users WHERE ?ids",
			params: map[string]any{
				"ids": []string{},
			},
			wantQuery: "SELECT * FROM users WHERE 1=1",
			wantArgs:  []any{},
		},
		{
			name:  "Slice int empty",
			query: "SELECT * FROM users WHERE ?ids",
			params: map[string]any{
				"ids": []int{},
			},
			wantQuery: "SELECT * FROM users WHERE 1=1",
			wantArgs:  []any{},
		},
		{
			name:  "Slice any empty",
			query: "SELECT * FROM users WHERE ?ids",
			params: map[string]any{
				"ids": []any{},
			},
			wantQuery: "SELECT * FROM users WHERE 1=1",
			wantArgs:  []any{},
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
