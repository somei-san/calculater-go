package mathelement

import (
	"calculater/math_element/number"
	"calculater/math_element/operator"
)

type ElementType int

const (
	Number ElementType = iota
	Operator
)

func GetType(soruce string) (result ElementType, ok bool) {
	ok = true

	if _, e := number.Create(soruce); e == nil {
		result = Number
		return
	}

	if _, ok2 := operator.Create(soruce); ok2 {
		result = Operator
		return
	}

	// どれにも当てはまらなければ対象外
	ok = false
	return
}
