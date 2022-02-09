# Register Machine
Such machine is capable of executing a sequence of instructions with a set of storage 
elements called register.  

# Machine language
## reserved keywords
```
assign goto save restore reg label const op perform
```
# Remark
Data stored in registers is persistent among instructions, in other words, it make 
instructions stateless.
The control of flow is implemented by facility of `label`
Recursive process take advantage of a stack.