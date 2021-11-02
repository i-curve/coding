package coding

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
	}
	if err, ok := data.(error); ok && err != nil && err.Error() != "" {
		tmp = &coding{point.httpCode, point.code, err.Error(), point}
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
