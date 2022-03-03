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
	code1 := coding.New(0, coding.StatusOK, "123")
	if code1.Error() != "123" {
		t.Error("err error by 1")
	}
	if code1.Code() != 200 {
		t.Error("err code by 1")
	}
	code2 := coding.New(0, coding.StatusAccepted, "")
	if code2 != nil {
		t.Error("err message by 1")
	}
}

func Test2(t *testing.T) {
	var code coding.Code
	code = coding.New(0, coding.StatusNoContent, "xxx")
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
	code = coding.New(0, coding.StatusCreated, "created")
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
	code1 := coding.New(0, coding.StatusOK, "错误点1")
	code2 := coding.New(100, 0, "错误点2")
	code2.Wrap(code1)
	if !coding.Is(code2, code1) {
		t.Error("err is by 5")
	}
	code3 := coding.New(0, coding.StatusOK, "错误点3")
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
	code1 := coding.New(0, 100, "错误点1")
	code2 := coding.New(0, 200, "错误点2")
	code3 := coding.New(0, 300, "错误点3")
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
	code1 := coding.New(0, 100, "")
	if code1 != nil {
		t.Error("err New by 7")
	}
	var err error
	code1 = coding.New(0, 100, err)
	if code1 != nil {
		t.Error("err New by 7")
	}
	code1 = coding.New(0, 100, "错误点1")
	err = errors.New("错误点2")
	if code1 == nil {
		t.Error("err New by 7")
	}
	code1 = nil
	if code1 != nil {
		t.Error("err New by 7")
	}
	code1 = coding.New(0, 100, err)
	if code1 == nil {
		t.Error("err New by 7")
	}
	code1 = coding.New(0, 100, "错误点1")
	if code1 == nil {
		t.Error("err New by 7")
	}
	// --------------------
	code1 = coding.New(0, 100, "错误点1")
	if code1.Code() != 100 || code1.Message() != "错误点1" {
		t.Error("err New'value by 7")
	}
	code1 = coding.New(0, 200, err)
	if code1.Code() != 200 || code1.Message() != "错误点2" || code1.Message() != "错误点2" {
		t.Error("err New'Value by 7")
	}
}

func Test8(t *testing.T) {
	err := errors.New("错误点2")
	code1 := coding.New(100, 0, "错误点1")
	code2 := coding.New(200, 0, err)
	code3 := coding.New(300, 0, "错误点3")
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

func Test10(t *testing.T) {
	var err error
	var code coding.Code
	ans := coding.Append(code, err)
	if ans != nil {
		t.Error("err append by 10")
	}
	code1 := coding.New(100, 0, "错误点1")
	code2 := coding.Append(code1, err)
	if !coding.Is(code2, code1) {
		t.Error("err append by  10")
	}
	// printcode(code2)
	err = errors.New("错误点")
	// ans = coding.Append(code, 100, err)
	// printcode(ans)
	code2 = coding.Append(code1, err)
	if !coding.Is(code2, code1) {
		t.Error("err append by 10")
	}
	// printcode(code2)
}
func Test11(t *testing.T) {
	var str string
	var code coding.Code
	ans := coding.Append(code, str)
	if ans != nil {
		t.Error("err append by 11")
	}
	code1 := coding.New(0, 100, "错误点1")
	code2 := coding.Append(code1, str)
	if !coding.Is(code2, code1) {
		t.Error("err append by  11")
	}
	// printcode(code2)
	str = "错误点"
	// ans = coding.Append(code, 100, err)
	// printcode(ans)
	code2 = coding.Append(code1, str)
	if !coding.Is(code2, code1) {
		t.Error("err append by 11")
	}
	// printcode(code2)
}

func Test12(t *testing.T) {
	var code coding.Code
	code1 := coding.New(0, 100, "错误点1")
	code2 := coding.New(0, 200, "错误点2")
	code = coding.Append(code, "错误点")
	if code != nil {
		t.Error("err append by 12")
	}

	if coding.Append(code1, "错误点11") == nil {
		t.Error("err append by 12")
	}
	if coding.Append(code2, "错误点22") == nil {
		t.Error("err append by 12")
	}
}

func Test13(t *testing.T) {
	var err error
	code := coding.New(0, 100, "错误点1")
	err = code
	type iCode interface {
		HTTPCode() int
		Message() string
		Code() int
	}
	var c iCode
	if errors.As(err, &c) {
		// fmt.Println("ok: ", c.Message(), c.Code(), c.HTTPCode())
	} else {
		t.Error("err As by 13")
	}
}

func Test14(t *testing.T) {
	var err coding.Code
	code := coding.New(0, 100, "错误点1")
	if coding.As(code, &err) {
		// fmt.Printf("%+v\n", err)
	} else {
		t.Error("err As by 14")
	}
}

func Test15(t *testing.T) {
	code1 := coding.New(0, 100, "错误点1")
	code2 := coding.New(0, 200, "错误点1")
	if coding.Is(code1, code2) {
		t.Error("err Is by 15")
	}
	if !coding.Is(code1, code1) {
		t.Error("err Is by 15")
	}
}
