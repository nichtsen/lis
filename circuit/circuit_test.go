package circuit

import (
	"fmt"
	"testing"
)

func TestHalfAdder(t *testing.T) {
	in := [2]*Wire{NewWire(), NewWire()}
	sum := NewWire()
	carry := NewWire()
	halfAdder := NewHalfAdder(in, sum, carry)

	in1 := []bool{true, false}
	in2 := []bool{true, false}
	for _, input1 := range in1 {
		for _, input2 := range in2 {
			in[0].SetSignal(input1)
			in[1].SetSignal(input2)
			fmt.Print(halfAdder)
		}
	}
}

func TestInverter(t *testing.T) {
	in := NewWire()
	out := NewWire()
	_ = NewInverter(in, out)
	in.SetSignal(true)
	fmt.Println(out.GetSignal())
	in.SetSignal(false)
	fmt.Println(out.GetSignal())
}
