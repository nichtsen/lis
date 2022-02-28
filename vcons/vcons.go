package vcons

import (
	"fmt"

	c "github.com/nichtsen/lis/cons"
)

type EType uint8
type EPointer uint64
type nullPtr uint64

const (
	Null    nullPtr = 0
	Pointer EType   = 1 << iota
	Content
)

type Element struct {
	ptrtyp  EType
	ptr     EPointer
	content interface{}
}

type Vector []*Element

var CarVector = make(Vector, 1024)
var CdrVector = make(Vector, 1024)
var free EPointer

func Cons(a, b interface{}) c.ICons {
	var ea, eb *Element

	defer func() {
		free++
	}()

	switch v := a.(type) {
	case EPointer:
		ea = &Element{
			ptrtyp: Pointer,
			ptr:    v,
		}
	default:
		ea = &Element{
			ptrtyp:  Content,
			content: a,
		}

	}

	switch v := b.(type) {
	case EPointer:
		eb = &Element{
			ptrtyp: Pointer,
			ptr:    v,
		}
	default:
		eb = &Element{
			ptrtyp:  Content,
			content: b,
		}

	}
	CarVector[free] = ea
	CdrVector[free] = eb
	return free
}

func Car(i interface{}) interface{} {
	if e, ok := i.(EPointer); !ok {
		panic("argument of car should be a EPointer")
	} else {
		return e.Car()
	}
}

func (e EPointer) Car() interface{} {
	ele := CarVector[e]
	switch ele.ptrtyp {
	case Pointer:
		return ele.ptr
	case Content:
		return ele.content
	default:
		panic("invalid element pointer type")

	}
}

func Cdr(i interface{}) interface{} {
	if p, ok := i.(EPointer); !ok {
		panic("argument of cdr should be a EPointer")
	} else {
		return p.Cdr()
	}
}

func (e EPointer) Cdr() interface{} {
	ele := CdrVector[e]
	switch ele.ptrtyp {
	case Pointer:
		return ele.ptr
	case Content:
		return ele.content
	default:
		panic("invalid element pointer type")

	}
}

func (e EPointer) SetCar(i interface{}) {
	var val *Element
	switch v := i.(type) {
	case EPointer:
		val = &Element{
			ptrtyp: Pointer,
			ptr:    v,
		}
	default:
		val = &Element{
			ptrtyp:  Content,
			content: i,
		}
	}
	CarVector[e] = val
}

func (e EPointer) SetCdr(i interface{}) {
	var val *Element
	switch v := i.(type) {
	case EPointer:
		val = &Element{
			ptrtyp: Pointer,
			ptr:    v,
		}
	default:
		val = &Element{
			ptrtyp:  Content,
			content: i,
		}
	}
	CdrVector[e] = val
}

func Pair(i interface{}) bool {
	_, ok := i.(EPointer)
	return ok
}

func Empty(i interface{}) bool {
	_, ok := i.(nullPtr)
	return ok
}

func List(args ...interface{}) c.ICons {
	if len(args) == 1 {
		return Cons(args[0], Null)
	}
	return Cons(args[0], List(args[1:]...))
}

func (n nullPtr) String() string {
	return "nil"
}

func (e EPointer) String() string {
	return fmt.Sprintf("(%v, %v)", e.Car(), e.Cdr())
}

func Map(i interface{}, callback func(i interface{})) {
	if _, ok := i.(nullPtr); ok {
		return
	}
	val := Car(i)
	callback(val)
	Map(Cdr(i), callback)
}
