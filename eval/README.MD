# eval
A tiny evaluator of scheme expressions  

## expression
- define
```
define { 
    var value-expr | 
    proc"(" [arg,] ")" 
    "{" 
         expr [, expr]
    "}" 
}
```
- assign
```
set var value-expr
```
- if
```
if predict-expr
    "{" 
    expr [, expr]
    "}"

```
- lambda
```
lambda([arg,]) { expr [,expr] }
```

- application
```
proc-name"("[,para]")" 
```

## Remark
Core of evaluator is interleaving of `eval` and `apply` loop. Everytime eval enconters a
application it call apply to make acctual calling.
A compound procedure is a sequence of expression to be evaluated 
within a given environment.

## Example 
compute fibonacci with evaluator:
```eval
Eval-go INPUT:
define fib(n) { 
    if ==(n,0) { 0 } 
    if ==(n,1) { 1 } 
    +(fib(-(n,1)),fib(-(n,2))) 
    };
Eval-go INPUT:
fib(10);
Eval-go OUTPUT:
 55 
```
