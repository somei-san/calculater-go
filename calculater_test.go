package main

import (
	"calculater/math_element/number"
	"testing"
)

/*正常系をチェックする*/
func exeNormalValidate(t *testing.T, args []string, expect int) {
	rslt, err := MainProc(args)

	if err != nil {
		t.Errorf("期待せぬエラーが発生した！ err = %v", err)
	}
	if rslt != number.Number(expect) {
		t.Errorf("%v を期待したのに %v だった！", expect, rslt)
	}
}

/*異常系をチェックする*/
func exeUnnormalValidate(t *testing.T, args []string) {
	expect := number.Number(0)
	rslt, err := MainProc(args)

	if rslt != expect {
		t.Errorf("%v を期待したのに %v だった！", expect, rslt)
	}
	if err == nil {
		t.Errorf("期待したエラーが発生しなかった！ err = %v", err)
	}
}

func Test_足し算(t *testing.T) {
	arg := []string{"12", "+", "12"}
	expect := 24

	exeNormalValidate(t, arg, expect)
}

func Test_足し算_結果マイナス(t *testing.T) {
	arg := []string{"12", "+", "-25"}
	expect := -13

	exeNormalValidate(t, arg, expect)
}

func Test_引き算(t *testing.T) {
	arg := []string{"25", "-", "12"}
	expect := 13

	exeNormalValidate(t, arg, expect)
}

func Test_引き算_結果マイナス(t *testing.T) {
	arg := []string{"12", "-", "25"}
	expect := -13

	exeNormalValidate(t, arg, expect)
}

func Test_掛け算(t *testing.T) {
	arg := []string{"12", "*", "7"}
	expect := 84

	exeNormalValidate(t, arg, expect)
}

func Test_掛け算_結果マイナス(t *testing.T) {
	arg := []string{"12", "*", "7"}
	expect := 84

	exeNormalValidate(t, arg, expect)
}

func Test_割り算_あまりなし(t *testing.T) {
	arg := []string{"12", "/", "4"}
	expect := 3

	exeNormalValidate(t, arg, expect)
}

func Test_割り算_あまりあり(t *testing.T) {
	arg := []string{"12", "/", "5"}
	expect := 2

	exeNormalValidate(t, arg, expect)
}

func Test_割り算_結果マイナス(t *testing.T) {
	arg := []string{"12", "/", "-5"}
	expect := -2

	exeNormalValidate(t, arg, expect)
}

func Test_複合(t *testing.T) {
	arg := []string{"2", "-", "3", "+", "4"}
	expect := 3
	exeNormalValidate(t, arg, expect)

	arg = []string{"2", "*", "3", "+", "4"}
	expect = 10
	exeNormalValidate(t, arg, expect)

	arg = []string{"2", "+", "3", "*", "4"}
	expect = 14
	exeNormalValidate(t, arg, expect)

	arg = []string{"2", "+", "3", "*", "4", "-", "4"}
	expect = 10
	exeNormalValidate(t, arg, expect)
}

func Test_数字や演算子以外の値を渡すとエラー(t *testing.T) {
	arg := []string{"あ", "*", "-5"}
	exeUnnormalValidate(t, arg)

	arg = []string{"12", "dfa", "-5"}
	exeUnnormalValidate(t, arg)

	arg = []string{"12", "*", "!"}
	exeUnnormalValidate(t, arg)
}

func Test_サイズ０の配列を渡すとエラー(t *testing.T) {
	arg := []string{}
	exeUnnormalValidate(t, arg)
}

func Test_演算子の位置がおかしい値を渡すとエラー(t *testing.T) {
	arg := []string{"-", "12", "*", "-5"}
	exeUnnormalValidate(t, arg)

	arg = []string{"12", "-", "*", "-5"}
	exeUnnormalValidate(t, arg)

	arg = []string{"12", "*", "-5", "-"}
	exeUnnormalValidate(t, arg)
}
