package main

import "fmt"

func MyFunc3() {
	println("hello world")
}

func MyFunc4() {
	myFunc3 := Func3
	str, err := myFunc3(1, 2)
	println(str, err)
}

func Func5() {
	fn := func(name string) string {
		return fmt.Sprintf("Hello, %s", name)
	}

	str := fn("Leo")

	println(str)
}
