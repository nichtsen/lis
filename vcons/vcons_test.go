package vcons

import (
	"fmt"
	"testing"
)

func TestList01(t *testing.T) {
	l := Cons(5, Cons(3, Cons(2, Cons(1, Cons(1, Null)))))
	ls := Cons("a", Cons("b", Cons("c", Null)))
	callbackl := func(i interface{}) {
		if n, ok := i.(int); !ok {
			t.Error("invalid content type")
		} else {
			fmt.Println(n)
		}

	}

	callbackls := func(i interface{}) {
		if s, ok := i.(string); !ok {
			t.Error("invalid content type")
		} else {
			fmt.Println(s)
		}

	}
	Map(l, callbackl)
	Map(ls, callbackls)
}

func ExampleList() {
	l := Cons(5, Cons(3, Cons(2, Cons(1, Cons(1, Null)))))
	ls := Cons("a", Cons("b", Cons("c", Null)))
	fmt.Print(l)
	fmt.Println("")
	fmt.Print(ls)
	//Output:
	// (5, (3, (2, (1, (1, nil)))))
	// (a, (b, (c, nil)))
}
