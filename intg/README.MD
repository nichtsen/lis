# intg
integral

## Remark
Use stream and delayed evaluation to compute integral 

## Stream
A stream is simply a sequence of objects. we can just represent it by iteration  

## Delayed Evaluation
Lambda expression is great for the purpose. 


## Benchmark
> solve a differential equation to estimate e
<div>
<img src="https://latex.codecogs.com/svg.image?y=e^x" title="e1" />
</div>
<div>
<img src="https://latex.codecogs.com/svg.image?f(y)=\frac{dy}{dx}" title="e2" />
</div>
<div>
<img src="https://latex.codecogs.com/svg.image?y(1)=\int_{0}^{1}f(y)dx&plus;y(0)" title="e3" />
</div>

`BS2` is implemented by Iterator
`BS3` is a Memoization version  
```shell
BenchmarkSolve/BS2-8         	       2	 827862354 ns/op
BenchmarkSolve/BS3-8         	     980	   1135534 ns/op
```
