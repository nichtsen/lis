package deriv

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	a := "日本語"
	fmt.Println(len(a))
	var r = []rune{'0', '1', '2', '9'}
	fmt.Println(r)
}

func TestInit(t *testing.T) {
	expr := "* 3 x"
	d := new(Differentiation)
	d.init(expr, 'x')
	fmt.Println(len(d.tokens))
	fmt.Println(string(d.mergeTokens()))
}

func ExampleDifferentiation() {
	expr := "* 3 x"
	expr1 := "+ (* 3 x) x"
	expr2 := "+ (* 3 x) (* x x)"
	d := new(Differentiation)
	res, err := d.Deriv(expr, 'x')
	if err != nil {
		fmt.Print(err)
		return
	}
	res2, err := d.Deriv(expr1, 'x')
	if err != nil {
		fmt.Print(err)
		return
	}
	res3, err := d.Deriv(expr2, 'x')
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println(res)
	fmt.Println(res2)
	fmt.Println(res3)
	// Output:
	// 3*1+0*x
	// 3*1+0*x+1
	// 3*1+0*x+x*1+1*x
}
