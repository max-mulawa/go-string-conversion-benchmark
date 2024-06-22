# String Cast Benchmarks for Golang

While coding Go I commonly find the need to do []byte(string) or string([]byte).


These simple benchmarks demystify the performance impacts of using those operations in your code.


## Usage
```go test -bench=.```

## My Results
On my Intel(R) Core(TM) i7-6820HQ CPU @ 2.70GHz I got these results, where using `unsafe` package yields best results.

```
goos: linux
goarch: amd64
pkg: string_cast_bench_test
cpu: Intel(R) Core(TM) i7-6820HQ CPU @ 2.70GHz
BenchmarkTestStringByteLen10-8                  49941799                24.48 ns/op
BenchmarkTestUnsafeStringByteLen10-8            1000000000               0.4852 ns/op
BenchmarkTestStringByteLen100-8                 29787290                43.09 ns/op
BenchmarkTestUnsafeStringByteLen100-8           1000000000               0.6580 ns/op
BenchmarkTestStringByteLen1000-8                 5422014               232.4 ns/op
BenchmarkTestStringByteLen10000-8                 805380              1417 ns/op
BenchmarkTestUnsafeStringByteLen10000-8         1000000000               0.4753 ns/op
BenchmarkTestByteStringLen10-8                  34550170                30.49 ns/op
BenchmarkTestUnsafeByteStringLen10-8            1000000000               0.6213 ns/op
BenchmarkTestCopyByteStringLen10-8              53925362                22.85 ns/op
BenchmarkTestByteStringLen100-8                 22402765                47.26 ns/op
BenchmarkTestUnsafeByteStringLen100-8           1000000000               0.6699 ns/op
BenchmarkTestCopyByteStringLen100-8             31352472                38.66 ns/op
BenchmarkTestByteStringLen1000-8                 5567014               213.0 ns/op
BenchmarkTestByteStringLen10000-8                 722914              1430 ns/op
BenchmarkTestUnsafeByteStringLen10000-8         1000000000               0.6480 ns/op
BenchmarkTestCopyByteStringLen10000-8             909958              1292 ns/op
PASS
ok      string_cast_bench_test  22.133s
```