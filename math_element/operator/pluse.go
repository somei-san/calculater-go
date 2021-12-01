package operator

import (
	"calculater/math_element/number"
)

type pluse struct {
}

func (pluse) Calc(x, y number.Number) number.Number {
	return number.Number(x + y)
}

func (pluse) IsHighPriority() bool {
	return false
}
