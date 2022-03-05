package eval

import (
	"fmt"
	"strings"
	"unicode"
)

type Expression []string

func MakeExpr(text string) *Expression {
	val := Expression(strings.Fields(strings.ReplaceAll(text, ";", " ")))
	return &val
}

func DefineExpr(expr Expression) bool { return expr[0] == "define" }
func AssignExpr(expr Expression) bool { return expr[0] == "set" }
func LambdaExpr(expr Expression) bool {
	return strings.HasPrefix(expr[0], "lambda") && ProcedureExpr(expr)
}
func IfExpr(expr Expression) bool     { return expr[0] == "if" }
func NumberExpr(expr Expression) bool { return strings.IndexFunc(expr[0], unicode.IsNumber) == 0 }
func SymbolExpr(expr Expression) bool { return strings.IndexRune(expr[0], '\'') == 0 }
func ApplicationExpr(expr Expression) bool {
	return strings.LastIndex(expr[0], ")") == len(expr[0])-1 && strings.IndexRune(expr[0], '(') > 0
}
func ApplicationaParas(expr Expression) []Expression {

	idx := strings.IndexRune(expr[0], '(')
	paraStr := strings.TrimSuffix(strings.TrimPrefix(expr[0][idx:], "("), ")")

	var args []string
	// inner Prodedure
	if strings.ContainsAny(paraStr, "()") {
		args = innerParas(strings.Split(paraStr, ","))
	} else {
		args = strings.Split(paraStr, ",")
	}

	res := make([]Expression, len(args))
	for idx, arg := range args {
		res[idx] = []string{arg}
	}
	return res
}

func innerParas(strs []string) []string {
	res := make([]string, 0)
	var count int
	var scan int
	for idx, str := range strs {
		if count > 0 {
			res[scan] = res[scan] + "," + strs[idx]
		}
		if count == 0 {
			res = append(res, strs[idx])
		}
		if count < 0 {
			panic(fmt.Sprintf("unclosed parenthese: %v", strs))
		}
		if strings.Index(str, "(") > 0 {
			if count == 0 {
				scan = len(res) - 1
			}
			count++
		}
		if strings.Index(str, ")") > 0 {
			count--
		}
	}
	return res
}

func ApplicationName(expr Expression) Expression {
	idx := strings.IndexRune(expr[0], '(')
	return Expression{expr[0][:idx]}
}
func ProcedureExpr(expr Expression) bool {
	return strings.LastIndex(expr[0], ")") == len(expr[0])-1 && strings.IndexRune(expr[0], '(') > 0 && expr[1] == "{"
}

func ProcedureVar(expr Expression) string {
	idx := strings.IndexRune(expr[0], '(')
	return expr[0][:idx]
}

func ProcedureParas(expr Expression) []string {
	idx := strings.IndexRune(expr[0], '(')
	return strings.Split(strings.Trim(expr[0][idx:], "()"), ",")
}

func ProcedureBody(expr Expression) (Expression, int) {
	val, idx := NextBlock(expr[1:])
	return val, idx + 2
}

func IfBody(expr Expression) (Expression, int) {
	val, idx := NextBlock(expr)
	return val, idx + 1
}

func (e Expression) Rest() Expression {
	return e[1:]
}

func (e Expression) Advance(n int) Expression {
	if n > len(e) {
		return e[len(e):]
	}
	return e[n:]
}

func NextBlock(e Expression) (Expression, int) {
	if e[0] != "{" {
		panic(fmt.Sprintf("Invalid block: %v", e))
	}
	var count = 1
	for i := 1; i < len(e); i++ {
		if e[i] == "}" {
			count--
			if count == 0 {
				return e[1:i], i
			}
		}
		if e[i] == "{" {
			count++
		}
	}
	panic(fmt.Sprintf("uneclosed block: %v", e))
}

type Procedure func(...interface{}) interface{}

func IsPrimitive(app interface{}) bool {
	_, ok := app.(Procedure)
	return ok
}

func Primitive(app interface{}) Procedure {
	return app.(Procedure)
}

type CompoundProcedure struct {
	paras []string
	body  Expression
	env   *Environment
}

func NewCompoundProdedure(paras []string, body Expression, env *Environment) CompoundProcedure {
	return CompoundProcedure{
		paras: paras,
		body:  body,
		env:   env,
	}
}

func IsCompound(app interface{}) bool {
	_, ok := app.(CompoundProcedure)
	return ok
}

func Compound(app interface{}) CompoundProcedure {
	return app.(CompoundProcedure)
}

func (c *CompoundProcedure) Body() Expression {
	return c.body
}

func (c *CompoundProcedure) Paras() []string {
	return c.paras
}

func (c *CompoundProcedure) Env() *Environment {
	return c.env
}
