package coding_test

import (
	"coding"
	"errors"
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	code1 := coding.New(coding.StatusOK, "123")
	if code1.Error() != "123" {
		t.Error("err error by 1")
	}
	if code1.Code() != 200 {
		t.Error("err code by 1")
	}
	code2 := coding.New(coding.StatusAccepted, "")
	if code2.Message() != "" {
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
	code2 := coding.New(coding.StatusOK, "错误点2")
	code2.Wrap(code1)
	// fmt.Println(coding.Unwrap(code2).Code())
	// fmt.Println(code1.Code())
	// fmt.Println(coding.Unwrap(code2).Code() == code1.Code())
	// fmt.Println(coding.Unwrap(code2).Error())
	// fmt.Println(code1.Error())
	if coding.Unwrap(code2).Code() != code1.Code() || coding.Unwrap(code2).Error() != code1.Error() {
		t.Error("err unwrap by 5")
	}
	if !coding.Is(code2, code1) {
		t.Error("err is by 5")
	}
	code3 := coding.New(coding.StatusOK, "错误点3")
	if !coding.Is(coding.Wrap(code3, code2), code1) {
		t.Error("err wrap by 5")
	}
	// fmt.Println(coding.Is(code2, code1))
}

func Test6(t *testing.T) {

}
