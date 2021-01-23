// Package rational data abstraction
package rational

import (
	"fmt"
	"strconv"
)

// Num interface of a retional number
type Num interface {
	Numerator() int
	Denominator() int
}

// RatNum representation of a rational number
type RatNum struct {
	numerator   int
	denominator int
}

// New return a retional num with numerator and denominator
func New(num int, den int) *RatNum {
	return &RatNum{
		numerator:   num,
		denominator: den,
	}
}

// Numerator return the numerator of a rational number
func (r *RatNum) Numerator() int {
	return r.numerator
}

// Denominator reutrn the denominator of a rational number
func (r *RatNum) Denominator() int {
	return r.denominator
}

// Float32 return float value of a rational number
func (r *RatNum) Float32() float32 {
	return float32(r.Numerator() / r.Denominator())
}

func (r *RatNum) String() string {
	return fmt.Sprintf("%s/%s", strconv.Itoa(r.numerator), strconv.Itoa(r.denominator))
}

func check(a Num, b Num) {
	if a.Denominator() == 0 || b.Denominator() == 0 {
		panic("Denominator can not be 0")
	}
}

// Add add
func Add(a Num, b Num) Num {
	check(a, b)
	return New(a.Numerator()*b.Denominator()+b.Numerator()*a.Denominator(), a.Denominator()*b.Denominator())
}

// Subtract subtract
func Subtract(a Num, b Num) Num {
	check(a, b)
	return New(a.Numerator()*b.Denominator()-b.Numerator()*a.Denominator(), a.Denominator()*b.Denominator())
}

// Multiply multiply
func Multiply(a Num, b Num) Num {
	check(a, b)
	return New(a.Numerator()*b.Numerator(), a.Denominator()*b.Denominator())
}

// Divide divide
func Divide(a Num, b Num) Num {
	check(a, b)
	return New(a.Numerator()*b.Denominator(), a.Denominator()/b.Numerator())
}

// Equal TODO
func Equal(a Num, b Num) bool {
	return false
}
