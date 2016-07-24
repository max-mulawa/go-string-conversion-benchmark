# String Cast Benchmarks for Golang

While coding Go I commonly find the need to do []byte(string) or string([]byte).


These simple benchmarks demystify the performance impacts of using those operations in your code.


## Usage
```go test -bench=.```

## My Results
On my i5-5200U CPU @ 2.20GHz I got these results

```
BenchmarkTestStringByteLen10-4          30000000                56.1 ns/op
BenchmarkTestStringByteLen100-4         20000000                92.9 ns/op
BenchmarkTestStringByteLen1000-4        20000000                95.5 ns/op
BenchmarkTestByteStringLen10-4          30000000                56.1 ns/op
BenchmarkTestByteStringLen100-4         20000000                92.4 ns/op
BenchmarkTestByteStringLen1000-4         3000000               389 ns/op
ok      github.com/ammario/go-string-cast-benchmark     11.007s
```