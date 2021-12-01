package formula

import (
	mathelement "calculater/math_element"
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
		s := source[index]
		_, ok := mathelement.GetType(source[index])
		if !ok {
			err = errors.New("解釈できない文字が入力されました。" + s)
			return
		}

		// ここまできたら演算子しかないはず
		operator, ok := operator.Create(s)
		if !ok {
			err = errors.New("バグ？" + s)
			return
		}

		index++
		if index >= len(source) {
			err = errors.New("入力フォーマットがヤバい。")
			return
		}

		operatorSource := source[index]
		num, e := number.Create(operatorSource)
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
