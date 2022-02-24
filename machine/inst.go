package machine

type Instruction struct {
	text Expression
	proc func()
}

type Expression []string
type Label string

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

func (i *Instruction) Expr() Expression {
	return i.text
}

func (e Expression) Rest() Expression {
	return e[1:]
}

func NewInstruction(text []string) *Instruction {
	return &Instruction{
		text: text,
	}
}

func AssignExpr(expr Expression) bool  { return expr.Tag() == "assign" }
func OpExpr(expr Expression) bool      { return expr.Tag() == "op" }
func TestExpr(expr Expression) bool    { return expr.Tag() == "test" && OpExpr(expr.Rest()) }
func BranchExpr(expr Expression) bool  { return expr.Tag() == "branch" }
func GotoExpr(expr Expression) bool    { return expr.Tag() == "goto" }
func LabelExpr(expr Expression) bool   { return expr.Tag() == "label" }
func RegExpr(expr Expression) bool     { return expr.Tag() == "reg" }
func NumberExpr(expr Expression) bool  { return expr.Tag() == "number" }
func StringExpr(expr Expression) bool  { return expr.Tag() == "string" }
func SaveExpr(expr Expression) bool    { return expr.Tag() == "save" }
func RestoreExpr(expr Expression) bool { return expr.Tag() == "restore" }
