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
