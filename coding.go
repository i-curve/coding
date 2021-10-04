package coding

import "strings"

type Code interface {
	Code() int
	Error() string
	Message() string
	Wrap(code Code) Code
	// Unwrap() Code
}

// coding .
type coding struct {
	code int
	text string
}

func (c coding) Code() int {
	return c.code
}
func (c coding) Error() string {
	return c.text
}
func (c coding) Message() string {
	if c.text == "" {
		return ""
	}
	return strings.Split(c.text, "\n")[0]
}

func (c *coding) Wrap(code Code) Code {
	c.code = code.Code()
	c.text = code.Error() + "\n" + c.text
	return c
}

// func (c *coding) Unwrap() Code {
// 	length := len(strings.Split(c.Error(), "\n"))
// 	if length <= 1 {
// 		return nil
// 	}
// 	c.text = strings.Join(strings.Split(c.Error(), "\n")[0:length-1], "\n")
// 	return c
// }

func New(code int, text string) Code {
	return &coding{code, text}
}
