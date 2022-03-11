package coding

import (
	"reflect"
)

// Append :append new node after the code.
// the first parameter is the code, the second parameter is the the message.It will be used when
// the code is not nil. So you can't judge if the code is nil.
//
// there is a decision.if the sourceFrom is nil, it will return a nil.
// the reasion is that I like Write code sequentially, and just due with error
// at the last.
func Append(sourceFrom interface{}, data interface{}) Code {
	if sourceFrom == nil {
		return nil
	}
	var (
		point *coding
		tmp   Code
	)
	sourceValue, okInit := sourceFrom.(*coding)
	if okInit {
		point = sourceValue
	}
	if text, ok := data.(string); ok && text != "" {
		tmp = &coding{point.httpCode, point.code, text, point}
	} else if text, ok := data.(*coding); ok && text != nil {
		tmp = &coding{text.httpCode, text.code, text.text, point}
	} else if err, ok := data.(error); ok && err != nil && err.Error() != "" {
		tmp = &coding{point.httpCode, point.code, err.Error(), point}
	}
	if tmp == nil && point != nil {
		tmp = point
	}
	return tmp
}

// Unwrap returns the result of calling the Unwrap method on err. If err's type
// isn't Coding.Code, Unwrap call Unwrap function. Otherwise, it returns err.
func Unwrap(code error) (result error) {
	if target, ok := code.(*coding); ok && target.point != nil {
		return target.point
	}
	if target, ok := code.(interface {
		Unwrap() error
	}); ok {
		return target.Unwrap()
	}
	return
}

// Is reports whether any data in target's chain matches target.
func Is(err, target error) bool {
	if target == nil || err == nil {
		return err == target
	}

	isCompareable := reflect.TypeOf(target).Comparable()
	for {
		if isCompareable && err == target {
			return true
		}
		if x, ok := err.(interface {
			Is(error) bool
		}); ok && x.Is(target) {
			return true
		}

		if err = Unwrap(err); err == nil {
			return false
		}
	}
}

func As(code error, target interface{}) bool {
	if target == nil {
		panic("errors: target cannot be nil")
	}
	val := reflect.ValueOf(target)
	typ := val.Type()
	if typ.Kind() != reflect.Ptr || val.IsNil() {
		panic("errors: target must be a pointer")
	}
	targetType := typ.Elem()
	if targetType.Kind() != reflect.Interface && !targetType.Implements(errorType) {
		panic("errors: target must be a pointer to an interface or implement error")
	}
	for code != nil {
		if reflect.TypeOf(code).AssignableTo(targetType) {
			val.Elem().Set(reflect.ValueOf(code))
			return true
		}
		if x, ok := code.(interface{ As(interface{}) bool }); ok && x.As(target) {
			return true
		}
		code = Unwrap(code)
	}
	return false
}

var errorType = reflect.TypeOf((*error)(nil)).Elem()
