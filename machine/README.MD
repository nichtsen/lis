# Simulated Register Machine
Such machine is capable of executing a sequence of instructions with a set of storage 
elements called register.  

# Machine language
## reserved keywords
```
assign save restore reg label test branch goto number string op perform
```
## 1. assign expression
assign reg-expr {primitive-value | op-expr} 
## 2. save expression 
save reg-key
## 3. restore expression
restore reg-key
## 4. reg expression
reg reg-key
## 5. label expression
label label-name
## 6. test expression
make a prediciton and set macnhine's flag register

test op-expr 
## 6. branch expression
branch label-expr
## 7. goto expression
goto {label-expr | reg-expr}
## 8. op expression
op operator-name {[,reg-expr] | [,number-expr] | [string-expr]}

# Stages
- assemble time
- simulate time

All instructions are evaluated during assemble time

# Remark
Data stored in registers is persistent among instructions, in other words, it make 
instructions stateless.
The control of flow is implemented by facility of `label`, which makes it an explicit way of controlling
Recursive process take advantage of a stack.
