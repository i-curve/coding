# Coding

![language](https://img.shields.io/badge/language-go-brightgreen) ![tests](https://img.shields.io/badge/tests-passing-brightgreen) [![License](https://img.shields.io/badge/license-BSD--3-brightgreen)](LICENSE)

参照 go 内置 errors 的实现方法实现 coding.Code.  
内部集成 string 和 error 这两种类型, 可以无缝切换使用.
用户向前端返回状态码和错误信息

## 具备功能

1. 可以像 error 一样通过 nil 进行判断改错误是否为空,不过不为空可以用 Error()方法获取错误信息
2. 具备链式功能,内部实现是一个链表
3. 内部具有状态有 HTTPCode,Code,Message,Error
4. 可以使用.Wrap()方法对错误进行包裹,链接不从层级错误
5. 可以使用 coding.Append()实现错误向后加累加的功能

## 实现原理

类似于 error 和 errors 的关系,首先定义了 Code 接口,并在内部用 code 实现的 Code 接口

定义的接口

```go
type Code interface {
	HTTPCode() int // http 返回状态码
	Code() int // 用户自定义状态码
	Error() string // 错误信息: 程序调试使用
	Message() string // 错误信息: 直接向用户展示部分
	Append(data interface{}) Code // 向错误信息后面添加数据
	Wrap(data interface{}) Code
	Unwrap() error
}
```

实现的结构体

```go
// coding .
type coding struct {
	httpCode int
	code     int
	text     string
	point    *coding
}
```

## 使用方法

```go
package main

import (
	"errors"
	"fmt"

	"github.com/i-curve/coding"
)

func main() {
	// coding.New 包方法
	// 第三个参数只有为string非""或者error非nil
	// 时 返回值才会非nil
	var err1 error
	err2 := errors.New("this is a error")
	code1 := coding.New(200, 0, "")   // code1 == nil
	code2 := coding.New(200, 0, err1) // code2 == nil
	code3 := coding.New(200, 0, err2) // code3 != nil
	fmt.Println("code1: ", code1, "code2: ", code2, "code3: ", code3)
	// code1:  <nil> code2:  <nil> code3:  this is a error

	// Append 包方法
	// 拼接错误 使用链表
	code1 = coding.New(500, 0, "Server Error")
	code2 = coding.New(404, 0, "err1")
	code3 = coding.Append(code1, code2)
	fmt.Println(code3.HTTPCode(), code3.Code(), code3.Message(), code3.Error())
	// 404 0 err1 err1;Server Error
	fmt.Println(coding.Is(code3, code1), coding.Is(code3, code2)) // true false
	fmt.Println(code3 == code1)                                   // false
	fmt.Println(code3.Error(), code1.Error(), code2.Error())
	// err1;Server Error Server Error err1
	coding.Append(nil, code3) // = nil

	// Append 成员方法
	code1 = coding.New(500, 0, "Server Error")
	code2 = coding.New(404, 0, "err1")
	code1.Append(code2)
	fmt.Println(code1.HTTPCode(), code1.Code(), code1.Message(), code1.Error())
	// 500 0 Server Error Server Error;err1
	code1.Append("err2")
	// Code{500, 0, "Server Error", "Server Error;err2"}
	err3 := errors.New("this is error3")
	code2.Append(err3)
	// Code{404, 0, "err1", "err1;this is error3"}

	// Unwrap 包方法
	var codex coding.Code
	err := coding.Unwrap(code3) // => code1 (error type)
	coding.As(err, &codex)      // => true
	fmt.Println(codex.HTTPCode(), codex.Code(), codex.Error())
	//  500 0 Server Error

	// Wrap 包方法
	// 添加一个新的阶段作为当前节点的子节点

	// As 包方法
	// 用于把Code类型的错误转换为指定的错误类型
	code1 = coding.New(500, 0, "Server Error")
	code2 = coding.New(404, 12, "err1")
	code3 = coding.Append(code1, code2)
	if coding.As(code3, &err) { // 返回是否转换成功
		fmt.Println("err:", err) // err: err1;Server Error
	}
	var code4 coding.Code
	if coding.As(code3, &code4) { // 返回是否转换成功
		fmt.Println(code4.HTTPCode(), code4.Code(), code4.Message(), code4.Error())
		// 404 12 err1 err1;Server Error
	}
}

```

## License

**[BSD-3-ClAUSE License](LICENSE)**
