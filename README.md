# go-hashmap-benchmark
Go hashmap read and write speed benchmark

```
goos: windows
goarch: amd64
pkg: benchmark-hashmap
BenchmarkWriteInt-8   	2000000000	         0.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkReadInt-8    	2000000000	         0.02 ns/op	       0 B/op	       0 allocs/op
BenchmarkWriteStr-8   	2000000000	         0.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkReadStr-8    	2000000000	         0.03 ns/op	       0 B/op	       0 allocs/op
PASS
coverage: 100.0% of statements
ok  	benchmark-hashmap	4.431s
Success: Benchmarks passed.
```