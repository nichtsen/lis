package eval

import (
	"fmt"
	"strconv"
)

func Eval(e *Expression, env *Environment) interface{} {
	for expr := (*e); len(expr) > 0; expr = (*e) {
		switch {
		case DefineExpr(expr):
			EvalDefine(e, env)
		case AssignExpr(expr):
			EvalAssign(e, env)
		case LambdaExpr(expr):
			continue
		case IfExpr(expr):
			val, ifContinue := EvalIf(e, env)
			if !ifContinue {
				return val
			}

		case NumberExpr(expr):
			n, _ := strconv.Atoi(expr[0])
			*e = expr[1:]
			return n
		case SymbolExpr(expr):
			return expr[0][1:]
		case ApplicationExpr(expr):
			ae := ApplicationName(expr)
			defer func() {
				*e = (*e)[1:]
			}()
			return Apply(Eval(&ae, env), EvalArgs(ApplicationaParas(expr), env)...)
		default:
			val := env.LookUpVariable(expr[0])
			*e = expr[1:]
			return val
		}
	}
	return nil
}

func EvalArgs(exprs []Expression, env *Environment) []interface{} {
	res := make([]interface{}, 0)
	for _, expr := range exprs {
		res = append(res, Eval(&expr, env))
	}
	return res
}

func Apply(app interface{}, args ...interface{}) interface{} {
	if IsPrimitive(app) {
		return Primitive(app)(args...)
	}
	if IsCompound(app) {
		cp := Compound(app)
		e := cp.Body()
		return Eval(&e, ExtendEnv(cp.Env(), cp.Paras(), args))
	}
	panic(fmt.Sprintf("Invalid application %v", app))
}

func EvalDefine(e *Expression, env *Environment) {
	(*e) = (*e)[1:]
	// define a procedure variable
	if ProcedureExpr((*e)) {
		expr := *e
		va := ProcedureVar(expr)
		paras := ProcedureParas(expr)
		body, idx := ProcedureBody(expr)
		cp := NewCompoundProdedure(paras, body, env)
		env.DefineVariable(va, cp)
		*e = (*e)[idx:]
		return

	}
	// define a non-proc variable
	va := (*e)[0]
	*e = (*e)[1:]
	env.DefineVariable(va, Eval(e, env))
}

func EvalAssign(e *Expression, env *Environment) {
	va := (*e)[1]
	*e = (*e)[2:]
	env.SetVariable(va, Eval(e, env))
}

func EvalIf(e *Expression, env *Environment) (interface{}, bool) {
	*e = (*e)[1:]
	val := Eval(e, env)
	v, ok := val.(bool)
	if !ok {
		panic(fmt.Sprintf("Invalid Prediction, %v", v))
	}
	csq, idx := IfBody(*e)
	defer func() {
		*e = (*e)[idx:]
	}()
	if v {
		val := Eval(&csq, env)
		if val != nil {
			return val, false
		}
	}
	return nil, true
}
