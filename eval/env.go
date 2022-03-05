package eval

import "fmt"

var GlobalEnv *Environment

func init() {
	InitGlobal()
}

func InitGlobal() {
	GlobalEnv = &Environment{
		Frame: Frame{
			"+":  Procedure(Add),
			"==": Procedure(Equal),
			">":  Procedure(Larger),
		},
		Enclose: nil,
	}
}

func Add(args ...interface{}) interface{} {
	return args[0].(int) + args[1].(int)
}

func Equal(args ...interface{}) interface{} {
	return args[0].(int) == args[1].(int)
}

func Larger(args ...interface{}) interface{} {
	return args[0].(int) > args[1].(int)
}

type Frame map[string]interface{}
type Environment struct {
	Enclose *Environment
	Frame   Frame
}

func NewFrame(vars []string, vals []interface{}) Frame {
	if len(vars) != len(vals) {
		panic("length of vars and vals should be equal")
	}
	res := make(Frame)
	for idx, v := range vars {
		res[v] = vals[idx]
	}
	return res
}

func ExtendEnv(enclose *Environment, vars []string, vals []interface{}) *Environment {
	return &Environment{
		Enclose: enclose,
		Frame:   NewFrame(vars, vals),
	}
}

func IsEmpty(env *Environment) bool {
	return env == nil
}

func (env *Environment) LookUpVariable(v string) interface{} {
	if val, ok := env.Frame[v]; ok {
		return val
	} else {
		if IsEmpty(env.Enclose) {
			panic(fmt.Sprintf("Undefined varible %v", v))
		} else {
			return env.Enclose.LookUpVariable(v)
		}
	}
}

func (env *Environment) SetVariable(v string, val interface{}) {
	if _, ok := env.Frame[v]; ok {
		env.Frame[v] = val
	} else {
		if IsEmpty(env.Enclose) {
			panic(fmt.Sprintf("Unbound varible %v", v))
		} else {
			env.Enclose.SetVariable(v, val)
		}
	}
}

func (env *Environment) DefineVariable(v string, val interface{}) {
	if _, ok := env.Frame[v]; ok {
		panic(fmt.Sprintf("varible has been defined %v", v))
	}
	env.Frame[v] = val
}
