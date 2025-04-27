# ewhere

[![Go Report Card](https://goreportcard.com/badge/github.com/AldiRvn/ewhere)](https://goreportcard.com/report/github.com/AldiRvn/ewhere)
[![Go Reference](https://pkg.go.dev/badge/github.com/AldiRvn/ewhere.svg)](https://pkg.go.dev/github.com/AldiRvn/ewhere)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

ewhere (easy where) is a Go package for dynamically parsing SQL where query with `?field` placeholders.

Automatically builds and cleans WHERE clauses based on parameter map input. Fast, flexible, and clean.

## 🎯 How It Works

- Write your SQL template using `?field` placeholders
- Call `ewhere.Parse(queryTemplate, params)`
- **If a param is missing, nil, or empty, the field will be automatically removed from the final query**
- Clean and safe output without leaving broken SQL or leftover conditions

Example:

```go
queryTemplate := "SELECT * FROM users WHERE ?name AND ?age"

params := map[string]interface{}{
	"name": "",
	"age":  30,
}

query, args := ewhere.Parse(queryTemplate, params)

fmt.Println("Query:", query)    //? Query: SELECT * FROM users WHERE age = ?
fmt.Println("Args:", args)      //? Args: [30]
```

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