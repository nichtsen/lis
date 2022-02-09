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

// Stack container of registers
type Stack struct {
	stack *list.List
}

func (s *Stack) Save(r *IRegister) {
	s.stack.PushBack(r)
}

func (s *Stack) Restore() *IRegister {
	e := s.stack.Front()
	return e.Value.(*IRegister)
}

// Machine is capable of executing a sequence of instructions
type Machine struct {
	registers map[string]IRegister
	stack     *list.List
	inst      []*Instruction
	pc        []*Instruction
	ops       map[string]func(...interface{}) interface{}
}

func NewMachine(regs []string, text [][]string, ops map[string]func(...interface{}) interface{}) *Machine {
	m := &Machine{
		stack: list.New(),
		pc:    make([]*Instruction, 0),
		inst:  make([]*Instruction, 0),
		ops:   ops,
	}
	regTab := make(map[string]IRegister)
	for _, r := range regs {
		regTab[r] = NewRegister(r)
	}
	m.registers = regTab
	m.inst = m.assemble(text)
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

func extractLabels(text [][]string, handler func([]*Instruction, map[string][]*Instruction) []*Instruction) []*Instruction {
	// exit of recursion
	if len(text) == 0 {
		return handler(make([]*Instruction, 0), make(map[string][]*Instruction))
	}

	cur := text[len(text)-1]

	rest := text[:len(text)-1]
	// label
	if len(cur) == 1 {
		return extractLabels(rest,
			func(insts []*Instruction, labels map[string][]*Instruction) []*Instruction {
				labels[cur[0]] = insts
				return handler(insts, labels)
			})
	} else {
		return extractLabels(rest,
			func(insts []*Instruction, labels map[string][]*Instruction) []*Instruction {
				return handler(append(insts, NewInstruction(cur)), labels)
			})
	}
}

func (m *Machine) updateInstruct(insts []*Instruction, labels map[string][]*Instruction) []*Instruction {
	for _, inst := range insts {
		inst.SetProc(m.makeProc(inst, labels))
	}
	return insts
}

func (m *Machine) makeProc(inst *Instruction, labels map[string][]*Instruction) func() {
	switch inst.Tag() {
	case "assign":
		args := inst.Rest()
		target := m.GetRegister(args[0])
		args = args.Rest()
		var valProc func() interface{}
		if args.Tag() == "op" {
			valProc = m.makeOperation(args)

		} else {
			valProc = m.makePrimitiveExpr(args)
		}
		return func() {
			target.Set(valProc())
		}
	default:
		panic("invalid expression")
	}
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
	case "reg":
		return func() interface{} {
			return m.GetRegisterContent(expr[1])
		}
	default:
		return nil
	}
}

func (m *Machine) AdvancePc() {
	m.pc = m.pc[1:]
}

func (m *Machine) Excute() {
	insts := m.pc
	insts[0].Run()
	if len(insts) == 1 {
		return
	}
	m.AdvancePc()
	m.Excute()
}

func (m *Machine) Start() {
	m.pc = m.inst
	m.Excute()
}
