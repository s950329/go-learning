package main

import "syntax/variables/demo"

func main() {
	var a int = 123
	println(a)
	var b = 123
	println(b)
	var c = 12.4
	println(c)

	println(demo.Global)
}

const (
	Status  = iota
	Status1 = iota
	Status2 = iota
)
