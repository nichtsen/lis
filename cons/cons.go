package cons

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidItemType = errors.New("invalid item type")
)

type Item interface {
	car() Item
	cdr() Item
	SetCar(interface{}) error
	SetCdr(interface{}) error
}

func Cons(a, b interface{}) Item {
	res := &item{
		composed: true,
	}
	if val, ok := a.(string); ok {
		res.a = &item{
			composed: false,
			val:      val,
		}
	} else if item, ok := a.(Item); ok {
		res.a = item
	} else {
		return nil
	}
	if b == nil {
		res.b = nil
		return res
	}

	if val, ok := b.(string); ok {
		res.b = &item{
			composed: false,
			val:      val,
		}
	} else if item, ok := b.(Item); ok {
		res.b = item
	} else {
		return nil
	}
	return res
}

func ConsList(a, b interface{}) Item {
	res := &lItem{
		composed: true,
	}
	if val, ok := a.(string); ok {
		res.a = &lItem{
			composed: false,
			val:      val,
		}
	} else if item, ok := a.(Item); ok {
		res.a = item
	} else {
		return nil
	}
	if b == nil {
		res.b = nil
		return res
	}

	if val, ok := b.(string); ok {
		res.b = &lItem{
			composed: false,
			val:      val,
		}
	} else if item, ok := b.(Item); ok {
		res.b = item
	} else {
		return nil
	}
	return res
}

func Car(i Item) Item {
	return i.car()
}

func Cdr(i Item) Item {
	return i.cdr()
}

type item struct {
	a        Item
	b        Item
	composed bool
	val      string
}

func (i *item) String() string {
	if i.composed {
		if i.b == nil {
			return fmt.Sprintf("(%v)", i.a)
		}
		return fmt.Sprintf("(%v, %v)", i.a, i.b)
	}
	return i.val
}

func (i *item) car() Item {
	return i.a
}

func (i *item) cdr() Item {
	return i.b
}

func (i *item) SetCar(val interface{}) error {
	if val, ok := val.(string); ok {
		i.a = &item{
			composed: false,
			val:      val,
		}
		return nil
	}
	if item, ok := val.(Item); ok {
		i.a = item
		return nil
	}
	return ErrInvalidItemType
}

func (i *item) SetCdr(val interface{}) error {
	if val, ok := val.(string); ok {
		i.a = &item{
			composed: false,
			val:      val,
		}
		return nil
	}
	if item, ok := val.(Item); ok {
		i.b = item
		return nil
	}
	return ErrInvalidItemType
}

type lItem struct {
	a        Item
	b        Item
	composed bool
	val      string
}

func (i *lItem) String() string {
	if !i.composed {
		return fmt.Sprintf("%v", i.val)
	}
	var str string
	i.traversals(&str)
	return fmt.Sprintf("(%s)", str)
}

func (i *lItem) traversals(str *string) {
	*str += fmt.Sprintf(" %v", i.a)

	if i.b != nil {
		if l, ok := i.b.(*lItem); ok {
			l.traversals(str)
		}
	}
}

func (i *lItem) car() Item {
	return i.a
}

func (i *lItem) cdr() Item {
	return i.b
}

func (i *lItem) SetCar(val interface{}) error {
	if val, ok := val.(string); ok {
		i.a = &lItem{
			composed: false,
			val:      val,
		}
		return nil
	}
	if item, ok := val.(Item); ok {
		i.a = item
		return nil
	}
	return ErrInvalidItemType
}

func (i *lItem) SetCdr(val interface{}) error {
	if val, ok := val.(string); ok {
		i.a = &lItem{
			composed: false,
			val:      val,
		}
		return nil
	}
	if item, ok := val.(Item); ok {
		i.b = item
		return nil
	}
	return ErrInvalidItemType
}

func List(params ...interface{}) Item {
	return list(0, params...)
}

func list(idx int, params ...interface{}) Item {
	// exit of recursion
	if idx == len(params)-1 {
		return ConsList(params[idx], nil)
	}
	return ConsList(params[idx], list(idx+1, params...))
}
