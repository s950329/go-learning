package main

import (
	"math"
	"unicode/utf8"
)

func main() {
	var a int = 456
	var b int = 123
	println(a + b)
	println(a - b)
	println(a * b)
	println(a / b)
	println(float64(a) / float64(b))

	println(math.Abs(-12.3))

	String()
	Byte()
	Bool()
}

func String() {
	println("hello")
	println(`hello`)
	println(len("hello"))
	println(len("hello妳好"))
	println(utf8.RuneCountInString("HELLO 你好"))
}

func Byte() {
	var a byte = 'a'
	println(a)

	var str string = "hello"
	var bs []byte = []byte(str)
	var str1 string = string(bs)
	println(str1)
}

func Bool() {
	var a bool = true
	var b bool = false
	println(a && b)
	println(a || b)

}
