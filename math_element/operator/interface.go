package operator

import (
	"calculater/math_element/number"
)

type Operator interface {
	Calc(x, y number.Number) number.Number

	IsHighPriority() bool
}
