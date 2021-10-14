package intg

import "fmt"

type Iterator func() (float64, Iterator)

func (it Iterator) Ref(n int) float64 {
	for i := 0; i < n; i++ {
		_, it = it()
	}
	res, _ := it()
	return res
}

func (it Iterator) PrintN(n int) {
	res := make([]float64, n)
	for i := 0; i < n; i++ {
		var val float64
		val, it = it()
		res[i] = val
	}
	fmt.Println(res)
}

func (it Iterator) Scale(multiply float64) Iterator {
	return func() (float64, Iterator) {
		val, next := it()
		return val * multiply, next.Scale(multiply)
	}
}

func (it Iterator) Add(it2 Iterator) Iterator {
	return func() (float64, Iterator) {
		val1, next1 := it()
		val2, next2 := it2()
		return val1 + val2, next1.Add(next2)
	}
}

func (it Iterator) Map(f func(float64) float64) Iterator {
	return func() (float64, Iterator) {
		val, next := it()
		return f(val), next.Map(f)
	}
}

func Integral(integrand Iterator, initial float64, dt float64) Iterator {
	return func() (float64, Iterator) {
		return initial, integral(initial, integrand, dt)
	}
}

func IntegralDelay(integrandDelay func() Iterator, initial float64, dt float64) Iterator {
	return func() (float64, Iterator) {
		return initial, func() (float64, Iterator) {
			itegrand := integrandDelay()
			val, iterFy := itegrand()
			res := initial + val*dt
			return res, integralDelay(res, iterFy, dt)
		}
	}
}

func integral(prev float64, integrand Iterator, dt float64) Iterator {
	val, iter := integrand()
	res := prev + val*dt
	// lambda expression as delayed evaluation
	return func() (float64, Iterator) {
		return res, integral(res, iter, dt)
	}
}

func integralDelay(prev float64, integrand Iterator, dt float64) Iterator {
	// lambda expression as delayed evaluation
	// all values are evaluated until calling of the delayed funciton
	return func() (float64, Iterator) {
		val, iter := integrand()
		res := prev + val*dt
		return res, integralDelay(res, iter, dt)
	}
}

// Delay data encapsulation is neccessary since taking address of function is not allowed
type Delay struct {
	call Iterator
}

func (d *Delay) Delay() Iterator {
	return d.call
}

// Solve differential equation f(y) = dy/dt
func Solve(f func(float64) float64, initial, dt float64) Iterator {
	fy := &Delay{}
	// delay the fy until it get feedback from y
	y := IntegralDelay(fy.Delay, initial, dt)
	fy.call = y.Map(f)
	return y
}

func IntegralDelay2(integrandDelay func() Iterator, initial float64, dt float64) Iterator {
	return func() (float64, Iterator) {
		return initial, func() Iterator {
			integrand := integrandDelay()
			integrand = integrand.Scale(dt)
			integrand = integrand.Add(IntegralDelay2(integrandDelay, initial, dt))
			return integrand
		}()
	}
}

func Solve2(f func(float64) float64, initial, dt float64) Iterator {
	var fy Iterator
	var delayFY = func() Iterator {
		return fy
	}
	// delay the fy until it get feedback from y
	y := IntegralDelay2(delayFY, initial, dt)
	fy = y.Map(f)
	return y
}

func IntegralDelay3(dp []float64, integrandDelay func() Iterator, initial float64, dt float64) Iterator {
	return func() (float64, Iterator) {
		return dp[0], func() Iterator {
			integrand := integrandDelay()
			return dpIter(dp[0], dp[1:], integrand, dt)
		}()
	}
}

func dpIter(prev float64, dp []float64, fy Iterator, dt float64) Iterator {
	return func() (float64, Iterator) {
		val, next := fy()
		dp[0] = prev + val*dt
		return dp[0], dpIter(dp[0], dp[1:], next, dt)
	}
}

// Memorization instead of computing from scratch
func FyMem(dp []float64, f func(float64) float64) Iterator {
	return func() (float64, Iterator) {
		return f(dp[0]), FyMem(dp[1:], f)
	}
}

func Solve3(f func(float64) float64, initial, dt float64) Iterator {
	var fy Iterator
	var delayFY = func() Iterator {
		return fy
	}
	dp := make([]float64, 1024*1024)
	dp[0] = initial
	// delay the fy until it get feedback from y
	y := IntegralDelay3(dp, delayFY, initial, dt)
	fy = FyMem(dp, f)
	return y
}
