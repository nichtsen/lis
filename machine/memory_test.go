package machine

import (
	"fmt"
	"testing"
)

func TestList01(t *testing.T) {
	l := Cons(5, Cons(3, Cons(2, Cons(1, Cons(1, Empty)))))
	ls := Cons("a", Cons("b", Cons("c", Empty)))
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
