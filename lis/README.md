# lis
lis

# benchmark
go test -v -bench .
```
goarch: amd64
pkg: lis
BenchmarkLIS-4          	13353184	        82.9 ns/op	      64 B/op	       1 allocs/op
BenchmarkLISdynamic-4   	11743406	       103 ns/op	      64 B/op	       1 allocs/op
PASS
ok  	lis	2.611s
```