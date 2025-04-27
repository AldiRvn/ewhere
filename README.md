# ewhere
ewhere (easy where) is a Go package for dynamically parsing SQL where query with ?field placeholders.

Automatically builds and cleans WHERE clauses based on parameter map input. Fast, flexible, and clean..

## 📦 Install ewhere

```bash
go get github.com/AldiRvn/ewhere
```

## ✨ Benchmark Summary

- ⏱️ Avg time per parse: **2550 ns** (~2.5 μs)
- 💾 Avg memory usage: **2630 bytes**
- 🔁 Avg allocations: **37 per operation**
- 🖥️ CPU: 12th Gen Intel(R) Core(TM) i5-12400F
- 📋 Benchmark runs: **469,850 iterations**

## 🧪 Benchmark Result

```bash
goos: windows
goarch: amd64
pkg: ewhere/ehwere
cpu: 12th Gen Intel(R) Core(TM) i5-12400F
=== RUN   BenchmarkParse
BenchmarkParse
BenchmarkParse-12
  469850              2550 ns/op            2630 B/op         37 allocs/op
PASS
ok      ewhere/ehwere   1.391s
```

## 🚀 Example Usage

```go
package main

import (
	"fmt"
	"github.com/AldiRvn/ewhere"
)

func main() {
	queryTemplate := "SELECT * FROM users WHERE ?name AND ?age"

	params := map[string]interface{}{
		"name": "Jane",
		"age":  25,
	}

	query, args := ewhere.Parse(queryTemplate, params)

	fmt.Println("Query:", query)  // Query: SELECT * FROM users WHERE name = ? AND age = ?
	fmt.Println("Args:", args)    // Args: [Jane 25]
}
```
