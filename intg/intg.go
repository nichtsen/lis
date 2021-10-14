package intg

type Iterator func() (float64, Iterator)

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
	fy.call = mapy(y, f)
	return y
}

func mapy(y Iterator, f func(float64) float64) Iterator {
	return func() (float64, Iterator) {
		res, next := y()
		// fmt.Println(res)
		return f(res), mapy(next, f)
	}
}
