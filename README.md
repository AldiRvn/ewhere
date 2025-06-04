# ewhere

[![Build Status](https://github.com/AldiRvn/ewhere/actions/workflows/coveralls.yml/badge.svg)](https://github.com/AldiRvn/ewhere/actions/workflows/coveralls.yml)
[![Coverage Status](https://coveralls.io/repos/github/AldiRvn/ewhere/badge.svg?branch=master)](https://coveralls.io/github/AldiRvn/ewhere?branch=master)
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

params := map[string]any{
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

- ⏱️ Avg time per parse: **2618 ns** (~2.6 μs)
- 💾 Avg memory usage: **981 bytes**
- 🔁 Avg allocations: **19 per operation**
- 🖥️ CPU: Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz
- 📋 Benchmark runs: **515,733 iterations**

## 🧪 Benchmark Result

```bash
goos: linux
goarch: amd64
pkg: github.com/AldiRvn/ewhere
cpu: Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz
BenchmarkParse
BenchmarkParse-5          515733              2618 ns/op             981 B/op         19 allocs/op
PASS
ok      github.com/AldiRvn/ewhere       1.382s
```
