package formula

import (
	"calculater/math_element/number"
	"calculater/math_element/operator"
)

type calcable interface {
	Calc() number.Number
}

type formula struct {
	Right    calcable
	Left     calcable
	CalcType operator.Operator
}

func (f formula) Calc() number.Number {
	return f.CalcType.Calc(f.Left.Calc(), f.Right.Calc())
}

func (f *formula) Insert(operator operator.Operator, num number.Number) {
	if operator.IsHighPriority() {
		rihteEnd := f.GetRightEnd()
		new := formula{
			Right:    num,
			CalcType: operator,
			Left:     rihteEnd.Right,
		}
		rihteEnd.Right = new
	} else {
		sub := *f
		new := formula{
			Right:    num,
			CalcType: operator,
			Left:     sub,
		}
		*f = new
		return
	}
}

func (f *formula) GetRightEnd() *formula {
	current := f
	for {
		if _, ok := current.Right.(number.Number); ok {
			return current
		}

		sub, ok := current.Right.(formula)
		if !ok {
			panic("マジムリ...")
		}
		current = &sub
	}

	return f
}
