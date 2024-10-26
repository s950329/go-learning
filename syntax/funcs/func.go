package main

func main() {
	Invoke()
}

func Invoke() {
	str := Func0("Leo")
	println(str)
	str1, err := Func2(12, 13)
	println(str1, err)
	_, err = Func3(1, 2)
	_, _ = Func1(1, 2, 3, "abc")
	Func1(1, 2, 3, "abc")

	Func5()
}

func Func0(name string) string {
	return "func0 " + name
}

func Func1(a, b, c int, d string) (string, error) {
	return "hello world", nil
}

func Func2(a int, b int) (str string, err error) {
	str = "hello world"
	return
}

func Func3(a int, b int) (str string, err error) {
	return "hello world", nil
}
