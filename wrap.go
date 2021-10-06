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
// func indirect(reflectValue reflect.Value) reflect.Value {
// 	for reflectValue.Kind() == reflect.Ptr {
// 		reflectValue = reflectValue.Elem()
// 	}
// 	return reflectValue
// }
// func indirectType(reflectType reflect.Type) (_ reflect.Type, isPtr bool) {
// 	for reflectType.Kind() == reflect.Ptr || reflectType.Kind() == reflect.Slice {
// 		reflectType = reflectType.Elem()
// 		isPtr = true
// 	}
// 	return reflectType, isPtr
// }
func Append(sourceFrom interface{}, code int, data interface{}) Code {
	var (
		point *coding
		tmp   Code
	)
	// sourceValue := indirect(reflect.ValueOf(sourceFrom))
	// sourceType, isPtr := indirectType(reflect.TypeOf(sourceFrom))
	// if sourceType == reflect.TypeOf(coding) {
	// 	point = &sourceValue
	// }
	sourceValue, okInit := sourceFrom.(*coding)
	if okInit {
		point = sourceValue
	}
	if text, ok := data.(string); ok && text != "" {
		tmp = &coding{code, text, point}
	}
	if err, ok := data.(error); ok && err != nil && err.Error() != "" {
		tmp = &coding{code, err.Error(), point}
	}
	if tmp == nil && point != nil {
		tmp = point
	}
	// if okInit {
	// 	sourceFrom = tmp
	// }
	return tmp
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
