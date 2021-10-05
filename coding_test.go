package coding_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/i-curve/coding"
)

func printcode(code coding.Code) {
	fmt.Println("print code:---------------")
	fmt.Println("code: ", code.Code())
	fmt.Println("message: ", code.Message())
	fmt.Println("error: ", code.Error())
	fmt.Println("end code_______________")
}
func Test1(t *testing.T) {
	code1 := coding.New(coding.StatusOK, "123")
	if code1.Error() != "123" {
		t.Error("err error by 1")
	}
	if code1.Code() != 200 {
		t.Error("err code by 1")
	}
	code2 := coding.New(coding.StatusAccepted, "")
	if code2 != nil {
		t.Error("err message by 1")
	}
}

func Test2(t *testing.T) {
	var code coding.Code
	code = coding.New(coding.StatusNoContent, "xxx")
	if code.Code() != coding.StatusNoContent {
		t.Error("err code by 2")
	}
	if code.Error() != "xxx" {
		t.Error("err error by 2")
	}
}

func Test3(t *testing.T) {
	var code coding.Code
	if code != nil {
		t.Error("err nil by 3")
	}
	code = coding.New(coding.StatusCreated, "created")
	if code == nil {
		t.Error("err nil by 3")
	}
	// fmt.Println(code)
}

func Test4(t *testing.T) {
	err := errors.New("xxx")
	err2 := fmt.Errorf("err2: %w", err)
	// fmt.Println(err)
	// fmt.Println(errors.Unwrap(err))
	if !errors.Is(err2, err) {
		t.Error("err by 4")
	}
	if errors.Is(err, err2) {
		t.Error("err by 4")
	}
}

func Test5(t *testing.T) {
	code1 := coding.New(coding.StatusOK, "错误点1")
	code2 := coding.New(100, "错误点2")
	code2.Wrap(code1)
	if !coding.Is(code2, code1) {
		t.Error("err is by 5")
	}
	code3 := coding.New(coding.StatusOK, "错误点3")
	if !coding.Is(code3.Wrap(code2), code1) {
		t.Error("err wrap by 5")
	}
	if !coding.Is(code3, code2) {
		t.Error("err wrap by 5")
	}
	// if coding.Is(code3.Wrap(code1), code2) {
	// 	t.Error("err wrap by 5")
	// }
	// if !coding.Is(code3, code1) {
	// 	t.Error("err wrap by 5")
	// }
	// code3.Wrap(code1)
	// fmt.Println("")
	// fmt.Println(code3.Error())
	// if !coding.Is(code3, code1) {
	// 	t.Error("err wrap by 5")
	// }

	// if !coding.Is(coding.Wrap(code3, code2), code1) {
	// 	t.Error("err wrap by 5")
	// }
	// fmt.Println(coding.Is(code2, code1))
}

func Test6(t *testing.T) {
	code1 := coding.New(100, "错误点1")
	code2 := coding.New(200, "错误点2")
	code3 := coding.New(300, "错误点3")
	code2.Wrap(code1)
	// if !coding.Is(code3.Wrap(code2), code1) {
	// 	t.Error("err wrap by 5")
	// }
	code3.Wrap(code2)
	if !coding.Is(code3, code1) {
		t.Error("err wrap by 6")
	}
	if !coding.Is(code3, code2) {
		t.Error("err wrap by 5")
	}

	code3.Wrap(code1)
	// printcode(coding.Unwrap(code3))
	// var data coding.Code
	// if data == nil {
	// 	fmt.Println("data == nil")
	// } else {
	// 	fmt.Println("data != nil")
	// }
	// data = nil
	// if data == nil {
	// 	fmt.Println("data == nil")
	// } else {
	// 	fmt.Println("data != nil")
	// }
	// fmt.Println("h: ", coding.Unwrap(code1))
	// if coding.Unwrap(code1) != nil {
	// 	t.Error("err unwrap by 6")
	// }
	// fmt.Println("test:", code1.Test(), ";", code1.Test() == nil)
	if coding.Unwrap(code1) != nil {
		t.Error("err unwrap by 6")
	}
	if coding.Is(code3, code2) {
		t.Error("err wrap by 5")
	}
	if coding.Is(code3.Wrap(code1), code2) {
		t.Error("err wrap by 5")
	}
	if !coding.Is(code3.Wrap(code1), code1) {
		t.Error("err wrap by 6")
	}
}

// -----------------------------
func Test7(t *testing.T) {
	code1 := coding.New(100, "")
	if code1 != nil {
		t.Error("err New by 7")
	}
	var err error
	code1 = coding.New(100, err)
	if code1 != nil {
		t.Error("err New by 7")
	}
	code1 = coding.New(100, "错误点1")
	err = errors.New("错误点2")
	if code1 == nil {
		t.Error("err New by 7")
	}
	code1 = nil
	if code1 != nil {
		t.Error("err New by 7")
	}
	code1 = coding.New(100, err)
	if code1 == nil {
		t.Error("err New by 7")
	}
	code1 = coding.New(100, "错误点1")
	if code1 == nil {
		t.Error("err New by 7")
	}
	// --------------------
	code1 = coding.New(100, "错误点1")
	if code1.Code() != 100 || code1.Message() != "错误点1" {
		t.Error("err New'value by 7")
	}
	code1 = coding.New(200, err)
	if code1.Code() != 200 || code1.Message() != "错误点2" || code1.Message() != "错误点2" {
		t.Error("err New'Value by 7")
	}
}

func Test8(t *testing.T) {
	err := errors.New("错误点2")
	code1 := coding.New(100, "错误点1")
	code2 := coding.New(200, err)
	code3 := coding.New(300, "错误点3")
	code2.Wrap(code1)
	code3.Wrap(code2)
	if coding.Unwrap(code1) != nil {
		t.Error("err unwrap by 8")
	}
	if coding.Unwrap(code2) != code1 {
		t.Error("err unwrap by 8")
	}
	if coding.Unwrap(code3) != code2 {
		t.Error("err unwrap by 8")
	}
	// ------------------
	if !coding.Is(code3, code1) || !coding.Is(code2, code1) || !coding.Is(code1, code1) {
		t.Error("err Is by 8")
	}
	if coding.Is(code1, code2) || coding.Is(code1, code3) || coding.Is(code2, code3) {
		t.Error("err Is by 8")
	}
	// fmt.Println(coding.Is(code2, code2))
	code3.Wrap(code1)
	if coding.Is(code3, code2) || !coding.Is(code3, code1) {
		t.Error("err wrap by 8")
	}
	if !coding.Is(code3.Unwrap(), code1) {
		t.Error("err unwrap by 8")
	}
}

// func Test9(t *testing.T) {
// 	err := errors.New("错误点2")
// 	code1 := coding.New(100, "错误点1")
// 	code2 := coding.New(200, err)
// 	code3 := coding.New(300, "错误点3")
// 	coding.Wrap(&code2, code1)
// 	coding.Wrap(&code3, code2)
// }
