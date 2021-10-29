# intg
integral

## Remark
Use stream and delayed evaluation to compute integral 

## Stream
A stream is simply a sequence of objects. we can just represent it by iteration  

## Delayed Evaluation
Lambda expression is great for the purpose. 


## Benchmark
`BS2` is implemented by Iterator
`BS3` is a Memoization version  
```shell
BenchmarkSolve/BS2-8         	       2	 827862354 ns/op
BenchmarkSolve/BS3-8         	     980	   1135534 ns/op
```
