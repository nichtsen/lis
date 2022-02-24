package machine

import (
	"container/list"
	"strconv"
)

// IRegister is the storage element of a machine
type IRegister interface {
	Set(interface{})
	Get() interface{}
	Key() string
}

type Register struct {
	key     string
	content interface{}
}

func NewRegister(k string) IRegister {
	return &Register{
		key: k,
	}
}

func (r *Register) Set(c interface{}) {
	r.content = c
}

func (r *Register) Get() interface{} {
	return r.content
}

func (r *Register) Key() string {
	return r.key
}

type Stack interface {
	Save(r IRegister)
	Restore(r IRegister)
	Init()
	Empty() bool
}

// Stack container of registers
type lStack struct {
	stack *list.List
}

func (s *lStack) Save(r IRegister) {
	s.stack.PushBack(r)
}

func (s *lStack) Restore() IRegister {
	e := s.stack.Front()
	return e.Value.(IRegister)
}

type mStack struct {
	ePointer interface{}
}

func (s *mStack) Save(r IRegister) {
	s.ePointer = Cons(r.Get(), s.ePointer)
}

func (s *mStack) Restore(r IRegister) {
	val := Car(s.ePointer)
	defer func() {
		s.ePointer = Cdr(s.ePointer)
	}()
	r.Set(val)
}

func (s *mStack) Init() {
	s.ePointer = Empty
}

func (s *mStack) Empty() bool {
	_, ok := s.ePointer.(emptyPtr)
	return ok
}

func NewStack() Stack {
	s := &mStack{}
	s.Init()
	return s
}

// Machine is capable of executing a sequence of instructions
type Machine struct {
	registers map[string]IRegister
	stack     Stack
	flag      IRegister
	insts     []*Instruction
	pc        []*Instruction
	ops       map[string]func(...interface{}) interface{}
}

func NewMachine(regs []string, text [][]string, ops map[string]func(...interface{}) interface{}) *Machine {
	m := &Machine{
		pc:    make([]*Instruction, 0),
		insts: make([]*Instruction, 0),
		ops:   ops,
		flag:  NewRegister("flag"),
		stack: NewStack(),
	}
	regTab := make(map[string]IRegister)
	for _, r := range regs {
		regTab[r] = NewRegister(r)
	}
	m.registers = regTab
	m.insts = m.assemble(text)
	return m
}

func (m *Machine) GetRegister(key string) IRegister {
	if r, ok := m.registers[key]; ok {
		return r
	} else {
		panic("register not exit")
	}
}

func (m *Machine) SetRegisterContent(key string, content interface{}) {
	if r, ok := m.registers[key]; ok {
		r.Set(content)
	} else {
		panic("register dose not exit")
	}
}

func (m *Machine) GetRegisterContent(key string) interface{} {
	if r, ok := m.registers[key]; ok {
		return r.Get()
	} else {
		panic("register dose not exit")
	}
}

func (m *Machine) GetOperation(key string) func(...interface{}) interface{} {
	if op, ok := m.ops[key]; ok {
		return op
	} else {
		panic("operation dose not exit")
	}
}

func (m *Machine) assemble(text [][]string) []*Instruction {
	return extractLabels(text, m.updateInstruct)
}

func extractLabels(text [][]string, handler func([]*Instruction, map[Label][]*Instruction) []*Instruction) []*Instruction {
	// exit of recursion
	if len(text) == 0 {
		return handler(make([]*Instruction, 0), make(map[Label][]*Instruction))
	}

	cur := text[0]

	rest := text[1:]
	// label
	if len(cur) == 1 {
		return extractLabels(rest,
			func(insts []*Instruction, labels map[Label][]*Instruction) []*Instruction {
				labels[Label(cur[0])] = insts
				return handler(insts, labels)
			})
	} else {
		return extractLabels(rest,
			func(insts []*Instruction, labels map[Label][]*Instruction) []*Instruction {
				return handler(append(insts, NewInstruction(cur)), labels)
			})
	}
}

func (m *Machine) updateInstruct(insts []*Instruction, labels map[Label][]*Instruction) []*Instruction {
	for _, inst := range insts {
		inst.SetProc(m.makeProc(inst, labels))
	}
	return insts
}

func (m *Machine) makeProc(inst *Instruction, labels map[Label][]*Instruction) func() {
	expr := inst.Expr()
	switch {
	case AssignExpr(expr):
		args := expr.Rest()
		target := m.GetRegister(args[0])
		args = args.Rest()
		var valProc func() interface{}
		if OpExpr(args) {
			valProc = m.makeOperation(args)

		} else {
			valProc = m.makePrimitiveExpr(args)
		}
		return func() {
			target.Set(valProc())
			m.AdvancePc()
		}
	case TestExpr(expr):
		args := expr.Rest()
		valProc := m.makeOperation(args)
		return func() {
			m.flag.Set(valProc())
			m.AdvancePc()
		}
	case BranchExpr(expr):
		return m.makeBranch(expr, labels)
	case GotoExpr(expr):
		return m.makeGoto(expr, labels)
	case SaveExpr(expr):
		return m.makeSave(expr)
	case RestoreExpr(expr):
		return m.makeRestore(expr)
	default:
		panic("invalid expression")
	}
}

func (m *Machine) makeSave(expr Expression) func() {
	arg := expr.Rest()
	reg := m.GetRegister(arg[0])
	return func() {
		m.stack.Save(reg)
		m.AdvancePc()
	}
}

func (m *Machine) makeRestore(expr Expression) func() {
	arg := expr.Rest()
	reg := m.GetRegister(arg[0])
	return func() {
		m.stack.Restore(reg)
		m.AdvancePc()
	}
}

func (m *Machine) makeGoto(expr Expression, labelTab map[Label][]*Instruction) func() {
	args := expr.Rest()
	if LabelExpr(args) {
		args = args.Rest()
		label := Label(args[0])
		if insts, ok := labelTab[label]; ok {
			return func() {
				m.pc = insts
			}
		}
		panic("invalid label")
	}
	if RegExpr(args) {
		args = args.Rest()
		reg := m.GetRegister(args[0])
		return func() {
			content := reg.Get()
			label, ok := content.(Label)
			if !ok {
				panic("register's content is not a label")
			}
			if insts, ok := labelTab[label]; ok {
				m.pc = insts
				return
			}
			panic("invalid label in register")
		}
	}
	panic("invalid goto expression ")
}

func (m *Machine) makeBranch(expr Expression, labelTab map[Label][]*Instruction) func() {
	args := expr.Rest()
	if LabelExpr(args) {
		args = args.Rest()
		label := Label(args[0])
		insts, ok := labelTab[label]
		if !ok {
			panic("invalid label")
		}

		return func() {
			content := m.flag.Get()
			flag, ok := content.(bool)
			if !ok {
				panic("flag register's content is not a boolean value")
			}
			if !flag {
				m.AdvancePc()
				return
			}
			m.pc = insts
		}
	}
	panic("invalid branch expression ")
}

func (m *Machine) makeOperation(expr Expression) func() interface{} {
	op := m.GetOperation(expr[1])
	var aprocs []func() interface{}

	for i := 2; i < len(expr); i += 2 {
		aprocs = append(aprocs, m.makePrimitiveExpr(expr[i:i+2]))
	}
	return func() interface{} {
		var args []interface{}
		for _, aproc := range aprocs {
			args = append(args, aproc())
		}
		return op(args...)
	}
}

func (m *Machine) makePrimitiveExpr(expr Expression) func() interface{} {
	switch expr.Tag() {
	case "number":
		return func() interface{} {
			n, _ := strconv.Atoi(expr[1])
			return n
		}
	case "string":
		return func() interface{} {
			return expr[1]
		}
	case "reg":
		return func() interface{} {
			return m.GetRegisterContent(expr[1])
		}
	case "label":
		return func() interface{} {
			return Label(expr[1])
		}
	default:
		panic("invalid primitive expression")
	}
}

func (m *Machine) AdvancePc() {
	m.pc = m.pc[:len(m.pc)-1]
}

func (m *Machine) Excute() {
	insts := m.pc
	if len(insts) == 0 {
		return
	}
	insts[len(insts)-1].Run()
	m.Excute()
}

func (m *Machine) Start() {
	m.pc = m.insts
	m.Excute()
}
