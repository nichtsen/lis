package intg

type Iterator func() (float64, Iterator)

func Integral(integrand Iterator, initial float64, dt float64) Iterator {
	return func() (float64, Iterator) {
		return initial, integral(initial, integrand, dt)
	}
}

func integral(prev float64, integrand Iterator, dt float64) Iterator {
	val, iter := integrand()
	res := prev + val*dt
	// lambda expression as deferred evaluation
	return func() (float64, Iterator) {
		return res, integral(res, iter, dt)
	}
}

// sove differential equation f(y) = dy/dt
func Solve(f func(float64) float64, initial, dt float64) Iterator {
	var fy Iterator
	y := Integral(fy, initial, dt)
	fy = fc(y, f)
	return y
}

func fc(iter Iterator, f func(float64) float64) Iterator {
	res, next := iter()
	return func() (float64, Iterator) {
		return f(res), fc(next, f)
	}
}
