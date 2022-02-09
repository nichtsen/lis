package machine

type Instruction struct {
	text Expression
	proc func()
}

type Expression []string

func (i *Instruction) Run() {
	i.proc()
}

func (i *Instruction) SetProc(proc func()) {
	i.proc = proc
}

func (i *Instruction) Tag() string {
	return i.text.Tag()
}
func (e Expression) Tag() string {
	return e[0]
}

func (i *Instruction) Rest() Expression {
	return i.text.Rest()
}

func (e Expression) Rest() Expression {
	return e[1:]
}

func NewInstruction(text []string) *Instruction {
	return &Instruction{
		text: text,
	}
}
