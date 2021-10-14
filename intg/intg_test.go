package intg

import (
	"fmt"
	"math"
	"testing"
)

// Test function y=x on interval [0, 1]
func TestIntegral(t *testing.T) {
	var initial float64 = 0
	var dt float64 = 0.001
	iter := func() (float64, Iterator) {
		return initial, func() (float64, Iterator) {
			return iteration(initial, dt)
		}
	}
	intg := Integral(iter, 0, dt)
	cur := intg
	res := cur.Ref(1000)
	expected := 0.5
	diff := math.Abs(res - expected)
	fmt.Println(diff)
	if diff > 0.01 {
		t.Errorf("difference is larger than 0.01")
	}
}

// Test function f(x) = e^x and estimate e when x = 1
func TestE(t *testing.T) {
	// if y = e^x then dy/dt = f(y) = y
	f := func(y float64) float64 {
		return y
	}
	// with inital value of e^0 = 1
	var initial float64 = 1
	dt := 0.001
	var iter Iterator
	var res float64
	iter = Solve(f, initial, dt)
	res = iter.Ref(1000)
	fmt.Println(res)
}

func iteration(prev float64, dt float64) (float64, Iterator) {
	res := prev + dt
	return res, func() (float64, Iterator) {
		return iteration(res, dt)
	}
}
