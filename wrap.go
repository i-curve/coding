package coding

func Append(sourceFrom interface{}, code int, data interface{}) Code {
	var (
		point *coding
		tmp   Code
	)
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
