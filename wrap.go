package coding

// func Wrap(code *Code, data interface{}) Code {
// 	// if code == nil || reflect.TypeOf(code).Kind() != reflect.Ptr {
// 	// 	return nil
// 	// } else {
// 	// 	code :=
// 	// }
// 	codes, isCode := data.(Code)
// 	if isCode {
// 		target, isCode := codes.(*coding)
// 		if !isCode {
// 			ans, ok := (*code).(*coding)
// 			if code == nil || !ok {
// 				return target
// 			}
// 			ans.point = target
// 			return ans
// 		}
// 	}
// 	err, isError := data.(error)
// 	if isError {
// 		ans, ok := (*code).(*coding)
// 		if code == nil || !ok {
// 			return New(0, err.Error())
// 		}
// 		ans.point = &coding{0, err.Error(), nil}
// 		return ans
// 	}
// 	return *code
// }

func Append(source Code, code int, data interface{}) (result Code) {
	var point *coding
	if sourceValue, ok := source.(*coding); ok {
		point = sourceValue
	}
	if text, ok := data.(string); ok && text != "" {
		return &coding{code, text, point}
	}
	if err, ok := data.(error); ok && err != nil && err.Error() != "" {
		return &coding{code, err.Error(), point}
	}
	if point != nil {
		return point
	}
	return
}

func Unwrap(code Code) (result Code) {
	if target, ok := code.(*coding); ok && target.point != nil {
		return target.point
	}
	return
}

func Is(data, target Code) bool {
	if target == nil || data == nil {
		return data == target
	}
	for {
		if data.Code() == target.Code() && data.Message() == target.Message() && data.Error() == target.Error() {
			return true
		}

		if data = Unwrap(data); data == nil {
			return false
		}
	}
}

// TODO
// func As() {

// }
