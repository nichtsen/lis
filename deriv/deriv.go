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
	output   []*token
	valid    bool
	variable rune
}
type token struct {
	typ    int
	char   []rune
	number int
}

func (t *token) isNumber() bool {
	return t.typ == number_t
}

func (t *token) isVariable(variable rune) bool {
	if t.typ == symbol_t && t.char[0] == variable {
		return true
	}
	return false
}

func (d *Differentiation) init(expr string, variable rune) {
	d.valid = false
	d.tokens = make([]*token, 0)
	d.variable = variable

	expr = strings.ReplaceAll(expr, " ", "")
	expr = strings.ReplaceAll(expr, "(", "")
	expr = strings.ReplaceAll(expr, ")", "")

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
					char:   numberStk,
					number: num,
				}
				d.tokens = append(d.tokens, token)
				numberStk = numberStk[:0]
			}
			if tk == '+' {
				token := &token{
					typ:  sum_t,
					char: []rune{tk},
				}
				d.tokens = append(d.tokens, token)
				continue
			}
			if tk == '*' {
				token := &token{
					typ:  product_t,
					char: []rune{tk},
				}
				d.tokens = append(d.tokens, token)
				continue
			}
			token := &token{
				typ:  symbol_t,
				char: []rune{tk},
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
	d.output = d.deriv(d.tokens)
	return string(d.merge()), nil
}

func (d *Differentiation) deriv(toks []*token) []*token {
	tok := toks[0]
	switch tok.typ {
	case operand_t & tok.typ:
		if tok.isVariable(d.variable) {
			return []*token{
				&token{
					typ:    number_t,
					char:   []rune{'1'},
					number: 1,
				},
			}
		}
		return []*token{
			&token{
				typ:    number_t,
				char:   []rune{'0'},
				number: 0,
			},
		}
	case operator_t & tok.typ:
		opToks := toks[1:]
		switch tok.typ {
		case sum_t:
			addend, rem := d.next(opToks, []rune{})
			augend, _ := d.next(rem, []rune{})
			return d.MakeSum(d.deriv(addend), d.deriv(augend))
		case product_t:
			multiplier, rem := d.next(opToks, []rune{})
			multiplicand, _ := d.next(rem, []rune{})
			return d.MakeSum(
				d.MakeProduct(multiplier, d.deriv(multiplicand)),
				d.MakeProduct(d.deriv(multiplier), multiplicand),
			)
		}
	}
	return []*token{}
}

// next get next token recursively
// eg. next((+(* 3 x) 5) (* a x))
// returns [+ * 3 x 5]
func (d *Differentiation) next(toks []*token, opSatck []rune) ([]*token, []*token) {
	tok := toks[0]
	res := []*token{
		tok,
	}
	switch tok.typ {
	case operand_t & tok.typ:
		// exit of recursion
		if len(opSatck) == 0 || len(toks) == 0 {
			return res, toks[1:]
		}
		more, rem := d.next(toks[1:], opSatck[1:])
		res = append(res, more...)
		return res, rem
	case operator_t & tok.typ:
		opSatck = append(opSatck, tok.char[0])
		more, rem := d.next(toks[1:], opSatck)
		res = append(res, more...)
		return res, rem
	}
	return []*token{}, []*token{}
}

func (d *Differentiation) merge() []rune {
	res := []rune{}
	for _, tok := range d.output {
		res = append(res, tok.char...)
	}
	return res
}

func (d *Differentiation) mergeTokens() []rune {
	res := []rune{}
	for _, tok := range d.tokens {
		res = append(res, tok.char...)
	}
	return res
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
	res := make([]*token, 0, len(exprA)+len(exprB)+1)
	res = append(res, exprA...)
	res = append(res,
		&token{
			typ:  product_t,
			char: []rune{'+'},
		},
	)
	res = append(res, exprB...)
	return res
}

func (d *Differentiation) MakeProduct(exprA, exprB []*token) []*token {
	res := make([]*token, 0, len(exprA)+len(exprB)+1)
	res = append(res, exprA...)
	res = append(res,
		&token{
			typ:  product_t,
			char: []rune{'*'},
		},
	)
	res = append(res, exprB...)
	return res
}
