package main

type Outer struct {
	Inner
}

type Outer1 struct {
	*Inner
}

func (i Outer1) Name() string {
	return "outer"
}

type Inner struct {
}

func (i Inner) Name() string {
	return "inner"
}

func (i Inner) Hello() {
	println("Hello im ", i.Name())
}

func Components() {
	var o Outer
	o.Hello()

	o1 := Outer1{
		Inner: &Inner{},
	}
	o1.Hello()
}
