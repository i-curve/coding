# Coding

参照go内置errors的实现方法实现coding.Code
用户向前端返回状态码和错误信息

## 具备功能

1. 可以像error一样通过nil进行判断改错误是否为空,不过不为空可以用Error()方法获取错误信息
2. 具备链式功能,内部实现是一个链表
3. 可以使用.Wrap()方法对错误进行包裹,链接不从层级错误
4. 可以使用coding.Append()实现错误向后加累加的功能

## 实现原理

类似于error和errors的关系,首先定义了Code接口,并在内部用code实现的Code接口

定义的接口
```go
type Code interface {
	HTTPCode() int
	Code() int
	Error() string
	Message() string
	Wrap(data interface{}) Code
	Unwrap() Code
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

1. 可以用时候为nil判断是否有错误信息

```go
// todo
```

2. 内置方法

```go
// todo
```

## License

__LICENSE__