package operator

import (
	"calculater/math_element/number"
)

type divide struct {
}

func (divide) Calc(x, y number.Number) number.Number {
	return number.Number(x / y)
}

func (divide) IsHighPriority() bool {
	return true
}
