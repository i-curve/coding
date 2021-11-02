package coding

type Code interface {
	HTTPCode() int
	Code() int
	Error() string
	Message() string
	Wrap(data interface{}) Code
	Unwrap() Code
}

// coding .
type coding struct {
	httpCode int
	code     int
	text     string
	point    *coding
}

// HTTPCode return the http status code for front end.
func (c *coding) HTTPCode() int {
	return c.httpCode
}

// Code return the http status code for front end.
func (c coding) Code() int {
	return c.code
}

// Error return the error message for coder.
func (c *coding) Error() (str string) {
	point := c
	for {
		if point == nil {
			break
		}
		if str != "" {
			str += ";" + point.text
		} else {
			str += point.text
		}
		point = point.point
	}
	return str
}

// Message return the error message for users.
func (c coding) Message() string {
	return c.text
}

func (c *coding) Wrap(data interface{}) Code {
	code1, ok1 := data.(Code)
	if ok1 {
		if code, ok := code1.(*coding); ok {
			if c == nil {
				return code
			}
			c.point = code
			return c
		}
	}
	if code, ok2 := data.(error); ok2 && !ok1 {
		if c == nil {
			return New(0, 0, code.Error())
		}
		c.point = &coding{0, 0, code.Error(), nil}
	}
	return c
}

func (c *coding) Unwrap() (result Code) {
	if c == nil || c.point == nil {
		return
	}
	return c.point
}

// New accept two args:code and data. data only accepts string or coding.Code,otherwise, it
// will return nil.
// Note: if the string is "" or error.Error() equal "", it alse return nil.
func New(httpCode, code int, data interface{}) Code {
	if text, ok := data.(string); ok && text != "" {
		return &coding{httpCode, code, text, nil}
	}
	if err, ok := data.(error); ok && err != nil && err.Error() != "" {
		return &coding{httpCode, code, err.Error(), nil}
	}
	return nil
}
