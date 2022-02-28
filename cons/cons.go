package cons

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidItemType = errors.New("invalid item type")
)

type ICons interface {
	Car() interface{}
	Cdr() interface{}
	SetCar(interface{})
	SetCdr(interface{})
}

func Cons(a, b interface{}) ICons {
	return &item{
		a: a,
		b: b,
	}
}

func Car(i ICons) interface{} {
	return i.Car()
}

func Cdr(i ICons) interface{} {
	return i.Cdr()
}

type item struct {
	a interface{}
	b interface{}
}

func (i *item) String() string {
	return fmt.Sprintf("(%v, %v)", i.a, i.b)
}

func (i *item) Car() interface{} {
	return i.a
}

func (i *item) Cdr() interface{} {
	return i.b
}

func (i *item) SetCar(val interface{}) {
	i.a = val
}

func (i *item) SetCdr(val interface{}) {
	i.b = val
}

func List(args ...interface{}) ICons {
	if len(args) == 1 {
		return Cons(args[0], nil)
	}
	return Cons(args[0], List(args[1:]...))
}
