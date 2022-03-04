# eval
evaluator of scheme expressions 

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
- sequence
- variable
- lambda
- application
```
proc-name"("[,para]")" 
```

## Remark
Core of evaluator is interleaving of `eval` and `apply` loop. Everytime eval enconters a
application it call apply to make acctual calling.
A compound procedure is a sequence of expression to be evaluated 
within a given environment.