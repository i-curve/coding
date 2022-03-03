package coding

type Code interface {
	HTTPCode() int
	Code() int
	Error() string
	Message() string
	Append(data interface{}) Code
	Wrap(data interface{}) Code
	Unwrap() error
}

// coding .
type coding struct {
	httpCode int
	code     int
	text     string
	point    *coding
}

// HTTPCode return the http status code for front end.
func (c *coding) HTTPCode() (result int) {
	if c != nil {
		result = c.httpCode
	}
	return
}

// Code return the http status code for front end.
func (c *coding) Code() (result int) {
	if c != nil {
		result = c.code
	}
	return
}

// Message return the error message for users.
func (c *coding) Message() string {
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

func (c *coding) Unwrap() (result error) {
	if c == nil || c.point == nil {
		return
	}
	return c.point
}

// New accept three args: httpCode, code and data. data only accepts string or error. otherwise, it
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

// Append next node in the c' next.
func (c *coding) Append(data interface{}) (code Code) {
	if c == nil {
		return c
	}
	if text, ok := data.(*coding); ok {
		c.point = text
	}
	if text, ok := data.(string); ok && text != "" {
		c.point = &coding{0, 0, text, nil}
	}
	if text, ok := data.(error); ok && text != nil && text.Error() != "" {
		c.point = &coding{0, 0, text.Error(), nil}
	}
	return c
}
