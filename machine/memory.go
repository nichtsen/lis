package machine

type EType uint8
type ePointer uint64
type emptyPtr uint64

const (
	Empty   emptyPtr = 0
	Pointer EType    = 1 << iota
	Content
)

type Element struct {
	ptrtyp  EType
	ptr     ePointer
	content interface{}
}

type Vector []*Element

var CarVector = make(Vector, 1024)
var CdrVector = make(Vector, 1024)
var free ePointer

func Cons(a, b interface{}) ePointer {
	var ea, eb *Element

	defer func() {
		free++
	}()

	switch v := a.(type) {
	case ePointer:
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
	case ePointer:
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
	if p, ok := i.(ePointer); !ok {
		panic("argument of car should be a ePointer")
	} else {
		e := CarVector[p]
		switch e.ptrtyp {
		case Pointer:
			return e.ptr
		case Content:
			return e.content
		default:
			panic("invalid element pointer type")

		}
	}
}

func Cdr(i interface{}) interface{} {
	if p, ok := i.(ePointer); !ok {
		panic("argument of cdr should be a ePointer")
	} else {
		e := CdrVector[p]
		switch e.ptrtyp {
		case Pointer:
			return e.ptr
		case Content:
			return e.content
		default:
			panic("invalid element pointer type")

		}
	}
}

func Pair(i interface{}) bool {
	_, ok := i.(ePointer)
	return ok
}

func Map(i interface{}, callback func(i interface{})) {
	if _, ok := i.(emptyPtr); ok {
		return
	}
	val := Car(i)
	callback(val)
	Map(Cdr(i), callback)
}
