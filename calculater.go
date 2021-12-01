package main

import (
	"calculater/math_element/formula"
	"calculater/math_element/number"
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	fmt.Println(flag.Args())
	args := flag.Args()

	rslt, err := MainProc(args)

	if err != nil {
		print_err(err)
	}

	fmt.Println("結果：", rslt)
}

// メイン処理
func MainProc(args []string) (result number.Number, err error) {

	formula, err := formula.Create(args)
	if err != nil {
		return
	}

	result = formula.Calc()
	return
}

func print_err(err error) {
	fmt.Println("処理できませんでした。", err)
}
