package cons

import (
	"fmt"
	"testing"
)

func TestCons(t *testing.T) {
	item := Cons("a", "b")

	item2 := Cons(Cons("a", "b"), "c")
	fmt.Println(item)
	fmt.Println(item2)
}

func TestList(t *testing.T) {
	l := List("a", "b", "c", "d")
	fmt.Println(l)
}

func TestListC(t *testing.T) {
	l := List(List("a", "b"), "d")
	fmt.Println(l)
}

func TestListL(t *testing.T) {
	l := List(List("a", "b"), List("c", "d"))
	fmt.Println(l)
}

func TestListB(t *testing.T) {
	l := List(Cons("a", "1"), Cons("b", "2"), Cons("c", "3"))
	fmt.Println(l)
}

func TestSet(t *testing.T) {
	item := Cons(Cons("a", "b"), "c")
	item.SetCar("a")
	fmt.Println(item)
}
