package circuit

import "fmt"

type HalfAdder struct {
	In     [2]*Wire
	Sum    *Wire
	Carry  *Wire
	or     *OrGate
	and    [2]*AndGate
	invert *Inverter
}

func NewHalfAdder(in [2]*Wire, sum, carry *Wire) *HalfAdder {
	ha := &HalfAdder{
		In:    in,
		Sum:   sum,
		Carry: carry,
	}

	w1 := NewWire()
	w2 := NewWire()

	ha.and[0] = NewAndGate(in, carry)
	ha.or = NewOrGate(in, w2)
	ha.invert = NewInverter(carry, w1)
	ha.and[1] = NewAndGate([2]*Wire{w1, w2}, sum)

	return ha
}

func (h *HalfAdder) String() string {
	return fmt.Sprintf("Input: %v, %v\nSum, Carry: %v, %v\n", h.In[0].GetSignal(), h.In[1].GetSignal(), h.Sum.GetSignal(), h.Carry.GetSignal())
}
