package circuit

type Gate struct {
	Ins  []*Wire
	Outs []*Wire
}

func NewGate(in []*Wire, out []*Wire) *Gate {
	return &Gate{
		Ins:  in,
		Outs: out,
	}
}

type Inverter struct {
	In  *Wire
	Out *Wire
}

func NewInverter(in, out *Wire) *Inverter {
	i := &Inverter{
		In:  in,
		Out: out,
	}
	i.In.AcceptAction(i.Process)
	return i
}

func (i *Inverter) Process() {
	new := i.In.GetSignal()
	i.Out.SetSignal(logicalNot(new))
}

type OrGate struct {
	In  [2]*Wire
	Out *Wire
}

func NewOrGate(in [2]*Wire, out *Wire) *OrGate {
	o := &OrGate{
		In:  in,
		Out: out,
	}
	o.In[0].AcceptAction(o.Process)
	o.In[1].AcceptAction(o.Process)
	return o
}

func (o *OrGate) Process() {
	newA := o.In[0].GetSignal()
	newB := o.In[1].GetSignal()
	o.Out.SetSignal(logicalOr(newA, newB))
}

type AndGate struct {
	In  [2]*Wire
	Out *Wire
}

func NewAndGate(in [2]*Wire, out *Wire) *AndGate {
	a := &AndGate{
		In:  in,
		Out: out,
	}
	a.In[0].AcceptAction(a.Process)
	a.In[1].AcceptAction(a.Process)
	return a
}

func (a *AndGate) Process() {
	newA := a.In[0].GetSignal()
	newB := a.In[1].GetSignal()
	a.Out.SetSignal(logicalAnd(newA, newB))
}

func logicalNot(signal bool) bool {
	return !signal
}

func logicalOr(a, b bool) bool {
	return a || b
}

func logicalAnd(a, b bool) bool {
	return a && b
}

func logicalXOr(a, b bool) bool {
	return ((!a) && b) || b && (!a)
}
