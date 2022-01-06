package formula

import (
	"calculater/math_element/number"
	"calculater/math_element/operator"

	"github.com/pkg/errors"
)

func Create(source []string) (instance formula, err error) {

	instance, source, e := createFirstNode(source)
	if e != nil {
		err = e
		return
	}

	index := 0
	for ; index < len(source); index++ {
		operatorSource := source[index]
		operator, ok := operator.Create(operatorSource)
		if !ok {
			err = errors.New("解釈できない文字が入力されました。" + operatorSource)
			return
		}

		index++
		if index >= len(source) {
			err = errors.New("入力フォーマットがヤバい。")
			return
		}

		numSource := source[index]
		num, e := number.Create(numSource)
		if e != nil {
			err = e
			return
		}

		instance.Insert(operator, num)
	}

	return

}

func createFirstNode(source []string) (node formula, prroceccedSoruce []string, err error) {

	if len(source) < 3 {
		err = errors.New("長さが足りません。" + string(len(source)))
		return
	}

	s := source[0]

	left, e := number.Create(s)
	if e != nil {
		err = errors.New("入力フォーマットがヤバい。" + s)
		return
	}

	operatorSource := source[1]
	operator, ok := operator.Create(operatorSource)
	if !ok {
		err = errors.New("入力フォーマットがヤバい。" + operatorSource)
		return
	}

	s = source[2]
	righte, e := number.Create(s)
	if e != nil {
		err = errors.New("入力フォーマットがヤバい。" + s)
		return
	}

	node = formula{
		Left:     left,
		CalcType: operator,
		Right:    righte,
	}

	prroceccedSoruce = source[3:]

	return
}
