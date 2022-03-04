package eval

import (
	"fmt"
	"reflect"
	"testing"
)

func ExampleEnvironment() {
	vars := []string{"a", "b", "c"}
	val := []interface{}{"a", 3.14, 1}

	env := ExtendEnv(GlobalEnv, vars, val)
	env.SetVariable("b", "b")
	env.DefineVariable("d", "d")
	vars = append(vars, "d")
	for _, va := range vars {
		val := env.LookUpVariable(va)
		fmt.Println(val)
	}
	// Output:
	// a
	// b
	// 1
	// d
}

func TestProdedureExpr(t *testing.T) {
	expr := Expression([]string{"foo(a,b,c)", "{", "set", "a", "1", "a+b+c", "}"})
	if !ProcedureExpr(expr) {
		t.Error("failed prediction of procedure expression")
	}
	va := ProcedureVar(expr)
	if va != "foo" {
		t.Errorf("expected \"foo\", got %v", va)
	}
	paras := ProcedureParas(expr)
	ep := []string{"a", "b", "c"}
	if !reflect.DeepEqual(paras, ep) {
		t.Errorf("expected to be equal, %v vs %v", paras, ep)
	}

	eb := expr[2:6]
	body, _ := ProcedureBody(expr)

	if !reflect.DeepEqual(body, eb) {
		t.Errorf("expected to be equal, %v vs %v", body, eb)
	}
}

func TestEval01(t *testing.T) {
	text := `define b 1 b`
	expr := MakeExpr(text)
	val := Eval(expr, GlobalEnv)
	if val != 1 {
		t.Errorf("expected to be 1, not %v", val)
	}
}

func TestEval02(t *testing.T) {
	text := `define a 1 define foo(v) { v } foo(a) `
	expr := MakeExpr(text)
	val := Eval(expr, GlobalEnv)
	if val != 1 {
		t.Errorf("expected to be 1, not %v", val)
	}
}
