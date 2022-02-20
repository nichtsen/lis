package machine

import (
	"fmt"
	"testing"
)

func TestMachine(t *testing.T) {
	add := func(args ...interface{}) interface{} {
		c := args[0].(int) + args[1].(int)
		_ = c
		return args[0].(int) + args[1].(int)

	}
	m := NewMachine(
		[]string{"a", "b", "c"},
		[][]string{{"assign", "a", "number", "1"}, {"assign", "b", "number", "2"}, {"assign", "c", "op", "+", "reg", "a", "reg", "b"}},
		map[string]func(args ...interface{}) interface{}{
			"+": add,
		},
	)
	m.Start()
	fmt.Print(m.GetRegisterContent("c"))
}

func TestStack(t *testing.T) {
	s := NewStack()
	r := NewRegister("r")
	r.Set(1)
	s.Save(r)
	r.Set(2)
	s.Save(r)
	r.Set(3)
	s.Save(r)
	r.Set("end")
	for !s.Empty() {
		s.Restore(r)
		fmt.Println(r.Get().(int))
	}
}

func TestBranch(t *testing.T) {
	equal := func(args ...interface{}) interface{} {
		return args[0].(int) == args[1].(int)

	}
	m := NewMachine(
		[]string{"a", "b", "val"},
		[][]string{
			{"assign", "a", "number", "1"},
			{"assign", "b", "number", "2"},
			{"test", "op", "==", "reg", "a", "reg", "b"},
			{"branch", "label", "next"},
			{"next"},
			{"assign", "val", "string", "next"},
		},
		map[string]func(args ...interface{}) interface{}{
			"==": equal,
		},
	)
	m.SetRegisterContent("val", "unsettled content")
	m.Start()
	fmt.Print(m.GetRegisterContent("val"))

	m2 := NewMachine(
		[]string{"a", "b", "val"},
		[][]string{
			{"assign", "a", "number", "1"},
			{"assign", "b", "number", "1"},
			{"test", "op", "==", "reg", "a", "reg", "b"},
			{"branch", "label", "next"},
			{"next"},
			{"assign", "val", "string", "next"},
		},
		map[string]func(args ...interface{}) interface{}{
			"==": equal,
		},
	)
	m2.SetRegisterContent("val", "unsettled content")
	m2.Start()
	fmt.Print(m2.GetRegisterContent("val"))

}
