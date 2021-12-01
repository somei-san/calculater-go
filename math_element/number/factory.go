package number

import (
	"strconv"

	"github.com/pkg/errors"
)

func Create(source string) (instance Number, err error) {
	instance = Number(0)
	i, e := strconv.Atoi(source)
	if e != nil {
		return instance, errors.New("数字が認識できない。" + source)
	}

	instance = Number(i)
	return
}
