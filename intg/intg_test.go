package intg

import (
	"fmt"
	"math"
	"testing"
)

func iteration(prev float64, dt float64) (float64, Iterator) {
	res := prev + dt
	return res, func() (float64, Iterator) {
		return iteration(res, dt)
	}
}

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

func ExampleScale() {
	var initial float64 = 0
	var dt float64 = 1
	var iter Iterator
	iter = func() (float64, Iterator) {
		return initial, func() (float64, Iterator) {
			return iteration(initial, dt)
		}
	}
	iter = iter.Scale(10)
	iter.PrintN(10)
	// Output:
	// [0 10 20 30 40 50 60 70 80 90]
}

func ExampleAdd() {
	var initial float64 = 0
	var initial2 float64 = 10
	var dt float64 = 1
	var iter, iter2 Iterator
	iter = func() (float64, Iterator) {
		return initial, func() (float64, Iterator) {
			return iteration(initial, dt)
		}
	}
	iter2 = func() (float64, Iterator) {
		return initial2, func() (float64, Iterator) {
			return iteration(initial2, dt)
		}
	}

	iter = iter.Add(iter2)

	iter.PrintN(10)
	//Output:
	//[10 12 14 16 18 20 22 24 26 28]
}

func ExampleMap() {
	var initial float64 = 0
	var dt float64 = 1
	var iter Iterator
	iter = func() (float64, Iterator) {
		return initial, func() (float64, Iterator) {
			return iteration(initial, dt)
		}
	}
	iter = iter.Map(func(x float64) float64 { return x * x })
	iter.PrintN(10)
	// Output:
	// [0 1 4 9 16 25 36 49 64 81]
}

// Test function f(x) = e^x and estimate e when x = 1
func TestSolve(t *testing.T) {
	// if y = e^x then dy/dt = f(y) = y
	f := func(y float64) float64 {
		return y
	}
	// with inital value of e^0 = 1
	var initial float64 = 1
	dt := 0.0001
	var iter Iterator
	var res float64
	iter = Solve(f, initial, dt)
	res = iter.Ref(10000)
	fmt.Println(res)
}

func TestSolve2(t *testing.T) {
	// if y = e^x then dy/dt = f(y) = y
	f := func(y float64) float64 {
		return y
	}
	// with inital value of e^0 = 1
	var initial float64 = 1
	dt := 0.05
	var iter Iterator
	var res float64
	iter = Solve2(f, initial, dt)
	res = iter.Ref(20)
	fmt.Println(res)
}

func TestSolve3(t *testing.T) {
	// if y = e^x then dy/dt = f(y) = y
	f := func(y float64) float64 {
		return y
	}
	// with inital value of e^0 = 1
	var initial float64 = 1
	dt := 0.001
	var iter Iterator
	var res float64
	iter = Solve3(f, initial, dt)
	res = iter.Ref(1000)
	fmt.Println(res)
}

func BenchmarkSolve(b *testing.B) {
	// if y = e^x then dy/dt = f(y) = y
	f := func(y float64) float64 {
		return y
	}
	// with inital value of e^0 = 1
	var initial float64 = 1
	dt := 0.05
	// bottleneck is addressed by Ref
	b.Run("BS2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			iter := Solve2(f, initial, dt)
			iter.Ref(20)
		}
	})
	b.Run("BS3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			iter := Solve3(f, initial, dt)
			iter.Ref(20)
		}
	})
}
