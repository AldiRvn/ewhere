package ewhere

import (
	"reflect"
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
			wantQuery: "SELECT * FROM users WHERE 1=1", // ✅ now expect 1=1
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotQuery, gotArgs := Parse(tt.query, tt.params)
			if gotQuery != tt.wantQuery {
				t.Errorf("got query = %v, want %v", gotQuery, tt.wantQuery)
			}
			if !reflect.DeepEqual(gotArgs, tt.wantArgs) {
				t.Errorf("got args = %v, want %v", gotArgs, tt.wantArgs)
			}
		})
	}
}
