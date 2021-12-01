package operator

import (
	"calculater/math_element/number"
)

type minus struct{}

func (minus) Calc(x, y number.Number) number.Number {
	return number.Number(x - y)
}

func (minus) IsHighPriority() bool {
	return false
}
