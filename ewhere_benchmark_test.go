package ewhere

import (
	"testing"
)

func BenchmarkParse(b *testing.B) {
	query := "SELECT * FROM users WHERE ?name AND ?age AND ?email AND ?country"
	params := map[string]interface{}{
		"name":    "Jane",
		"age":     25,
		"email":   "jane@example.com",
		"country": "Japan",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Parse(query, params)
	}
}
