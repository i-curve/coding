package coding

import (
	"reflect"
	"strings"
)

func Wrap(code, target Code) Code {
	return New(target.Code(), target.Error()+"\n"+code.Error())
}

func Unwrap(code Code) Code {
	length := len(strings.Split(code.Error(), "\n"))
	if length <= 1 {
		return nil
	}
	return New(code.Code(), strings.Join(strings.Split(code.Error(), "\n")[:length-1], "\n"))
}

func Is(code, target Code) bool {
	if target == nil {
		return code == target
	}
	isComparable := reflect.TypeOf(target).Comparable()
	for {
		if isComparable && code == target {
			return true
		}

		if code.Code() == target.Code() && code.Error() == target.Error() && code.Message() == target.Message() {
			return true
		}

		if code = Unwrap(code); code == nil {
			return false
		}
	}
}

// TODO
// func As() {

// }
