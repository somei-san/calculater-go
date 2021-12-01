package operator

import (
	"calculater/math_element/number"
)

type multiply struct {
}

func (multiply) Calc(x, y number.Number) number.Number {
	return number.Number(x * y)
}

func (multiply) IsHighPriority() bool {
	return true
}
