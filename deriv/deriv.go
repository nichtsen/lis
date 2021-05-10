package deriv

import (
	"errors"
	"strconv"
	"strings"
)

const (
	operand_t  = 0b0011
	operator_t = 0b1100
	number_t   = 0b0001
	symbol_t   = 0b0010
	sum_t      = 0b1000
	product_t  = 0b0100
)

type IDifferentiation interface {
	Deriv(expr string, variable rune) (string, error)
	IsNumber(token interface{}) bool
	IsVariable(token interface{}) bool
	MakeSum(exprA, exprB []*token) []*token
	MakeProduct(exprA, exprB []*token) []*token
}

type Differentiation struct {
	tokens   []*token
	valid    bool
	variable rune
}
type token struct {
	typ    int
	char   rune
	number int
}

func (t *token) isNumber() bool {
	return t.typ == number_t
}

func (t *token) isVariable(variable rune) bool {
	if t.typ == symbol_t && t.char == variable {
		return true
	}
	return false
}

func (d *Differentiation) init(expr string, variable rune) {
	d.valid = false
	d.tokens = make([]*token, 0)
	d.variable = variable

	strings.ReplaceAll(expr, " ", "")
	strings.ReplaceAll(expr, "(", "")
	strings.ReplaceAll(expr, ")", "")

	tks := []rune(expr)
	var numberStk []rune
	for _, tk := range tks {
		if tk >= '0' && tk <= '9' {
			numberStk = append(numberStk, tk)
			continue
		} else {
			if len(numberStk) > 0 {
				num, err := strconv.Atoi(string(numberStk))
				if err != nil {
					return
				}
				token := &token{
					typ:    number_t,
					number: num,
				}
				d.tokens = append(d.tokens, token)
				numberStk = numberStk[:0]
			}
			if tk == '+' || tk == '*' {
				token := &token{
					typ:  operator_t,
					char: tk,
				}
				d.tokens = append(d.tokens, token)
				continue
			}
			token := &token{
				typ:  operand_t,
				char: tk,
			}
			d.tokens = append(d.tokens, token)
		}
	}
	d.valid = true
}

func (d *Differentiation) Deriv(expr string, variable rune) (string, error) {
	d.init(expr, variable)
	if !d.valid {
		return "", errors.New("Invalid expression")
	}

}

func (d *Differentiation) deriv(toks []*token) []*token {
	tok := toks[0]
	switch tok.typ {
	case operand_t & tok.typ:
		if tok.isVariable(d.variable) {
			return []*token{
				&token{
					typ:    number_t,
					char:   '1',
					number: 1,
				},
			}
		}
		return []*token{
			&token{
				typ:    number_t,
				char:   '0',
				number: 0,
			},
		}
	case operator_t & tok.typ:
		switch tok.typ {
		case sum_t:
			addend, rem := d.getToken(toks[1:])
			augend, _ := d.getToken(rem)
			return d.MakeSum(d.deriv(addend), d.deriv(augend))
		case product_t:
			multiplier, rem := d.getToken(toks[1:])
			multiplicand, _ := d.getToken(rem)
			addend := d.MakeProduct(multiplier)
			return d.MakeSum()
			multiplyer, rem := d.getToken(toks[1:])
			augend, _ := d.getToken(rem)
			return d.MakeProduct()
		}
	}

}

func (d *Differentiation) SumNext(toks []*token) []*token {
	tok := toks[0]
	switch tok.typ {
	case tok.typ & operand_t:
		tmp := make([]*token, 0)
		tmp = append(tmp, tok)
		return d.deriv(tmp)
	case operator_t & tok.typ:
		return d.deriv(toks)
	}
	return make([]*token, 0)
}

func (d *Differentiation) IsNumber(tok interface{}) bool {
	if tk, ok := tok.(token); ok {
		return tk.isNumber()
	}
	return false
}

func (d *Differentiation) IsVariable(tok interface{}) bool {
	if tk, ok := tok.(token); ok {
		return tk.isVariable(d.variable)
	}
	return false
}

func (d *Differentiation) MakeSum(exprA, exprB []*token) []*token {
	return "TODO"
}

func (d *Differentiation) MakeProduct(exprA, exprB []*token) []*token {
	return "TODO"
}
