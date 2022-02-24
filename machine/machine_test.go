package machine

import (
	"fmt"
	"testing"
)

var (
	equal = func(args ...interface{}) interface{} {
		return args[0].(int) == args[1].(int)

	}
	add = func(args ...interface{}) interface{} {
		return args[0].(int) + args[1].(int)
	}
	multiply = func(args ...interface{}) interface{} {
		return args[0].(int) * args[1].(int)
	}
)

func TestAssign(t *testing.T) {
	m := NewMachine(
		[]string{"a", "b", "c"},
		[][]string{{"assign", "a", "number", "1"}, {"assign", "b", "number", "2"}, {"assign", "c", "op", "+", "reg", "a", "reg", "b"}},
		map[string]func(args ...interface{}) interface{}{
			"+": add,
		},
	)
	m.Start()
	if m.GetRegisterContent("c").(int) != 3 {
		t.Error("register A should store value number 2")
	}
}

func ExampleStack() {
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
		fmt.Print(r.Get().(int))
	}
	//Output:
	//321
}

func TestBranch(t *testing.T) {
	m := NewMachine(
		[]string{"a", "b", "val"},
		[][]string{
			{"assign", "a", "number", "1"},
			{"assign", "b", "number", "1"},
			{"test", "op", "==", "reg", "a", "reg", "b"},
			{"branch", "label", "next"},
			{"assign", "val", "op", "+", "number", "1", "reg", "val"},
			{"next"},
			{"assign", "val", "op", "+", "number", "2", "reg", "val"},
		},
		map[string]func(args ...interface{}) interface{}{
			"==": equal,
			"+":  add,
		},
	)
	m.SetRegisterContent("val", 0)
	m.Start()
	if v := m.GetRegisterContent("val").(int); v != 2 {
		t.Errorf("register val expected to store value number 2 not %d", v)

	}

	m2 := NewMachine(
		[]string{"a", "b", "val"},
		[][]string{
			{"assign", "a", "number", "1"},
			{"assign", "b", "number", "2"},
			{"test", "op", "==", "reg", "a", "reg", "b"},
			{"branch", "label", "next"},
			{"assign", "val", "op", "+", "number", "1", "reg", "val"},
			{"next"},
			{"assign", "val", "op", "+", "number", "2", "reg", "val"},
		},
		map[string]func(args ...interface{}) interface{}{
			"==": equal,
			"+":  add,
		},
	)
	m2.SetRegisterContent("val", 0)
	m2.Start()
	if v := m2.GetRegisterContent("val").(int); v != 3 {
		t.Errorf("register val expected to store value number 3 not %d", v)

	}

}

// TestGotoLabel goto a label expression
func TestGotoLabel(t *testing.T) {
	m := NewMachine(
		[]string{"a"},
		[][]string{
			{"assign", "a", "number", "0"},
			{"goto", "label", "next"},
			{"assign", "a", "op", "+", "number", "1", "reg", "a"},
			{"next"},
			{"assign", "a", "op", "+", "number", "2", "reg", "a"},
		},
		map[string]func(args ...interface{}) interface{}{
			"==": equal,
			"+":  add,
		},
	)
	m.Start()
	if m.GetRegisterContent("a").(int) != 2 {
		t.Error("register A should store value number 2")
	}
}

// TestGotoReg goto a register expression
func TestGotoReg(t *testing.T) {
	m := NewMachine(
		[]string{"a", "continue"},
		[][]string{
			{"assign", "a", "number", "0"},
			{"assign", "continue", "label", "next"},
			{"goto", "reg", "continue"},
			{"assign", "a", "op", "+", "number", "1", "reg", "a"},
			{"next"},
			{"assign", "a", "op", "+", "number", "2", "reg", "a"},
		},
		map[string]func(args ...interface{}) interface{}{
			"==": equal,
			"+":  add,
		},
	)
	m.Start()
	if m.GetRegisterContent("a").(int) != 2 {
		t.Error("register A should store value number 2")
	}
}

func TestStack(t *testing.T) {
	m := NewMachine(
		[]string{"a"},
		[][]string{
			{"assign", "a", "number", "0"},
			{"save", "a"},
			{"assign", "a", "number", "1"},
			{"restore", "a"},
		},
		map[string]func(args ...interface{}) interface{}{},
	)
	m.Start()
	if m.GetRegisterContent("a").(int) != 0 {
		t.Error("register A should store value number 0")
	}
}

func TestFactorial(t *testing.T) {
	m := NewMachine(
		[]string{"n", "val", "continue"},
		[][]string{
			{"save", "n"},
			{"assign", "continue", "label", "end"},
			{"save", "continue"},
			{"fact-loop"},
			{"test", "op", "==", "reg", "n", "number", "1"},
			{"branch", "label", "exit"},
			{"assign", "n", "op", "+", "reg", "n", "number", "-1"},
			{"assign", "continue", "label", "after-fact"},
			{"save", "n"},
			{"save", "continue"},
			// recursive process
			{"goto", "label", "fact-loop"},
			{"exit"},
			{"assign", "val", "number", "1"},
			{"goto", "label", "after-fact"},
			{"after-fact"},
			{"restore", "continue"},
			{"restore", "n"},
			{"assign", "val", "op", "*", "reg", "n", "reg", "val"},
			{"goto", "reg", "continue"},
			{"end"},
		},
		map[string]func(args ...interface{}) interface{}{
			"==": equal,
			"*":  multiply,
			"+":  add,
		},
	)
	m.SetRegisterContent("n", 5)
	m.Start()
	if v := m.GetRegisterContent("val").(int); v != 120 {
		t.Errorf("register A should store value number 120, not %d", v)
	}
}
