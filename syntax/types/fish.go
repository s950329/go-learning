package main

import "fmt"

type Fish struct {
}

func (f Fish) Swin() {

}

type FakeFish Fish

type Yu = Fish

func UseFish() {
	f1 := Fish{}
	f1.Swin()

	f2 := FakeFish{}
	//f2.Swin()

	f3 := FakeFish(f1)

	f1 = Fish(f2)
	fmt.Println(f1, f2, f3)

	yu := Yu{}
	yu.Swin()
}
