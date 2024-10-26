package main

import "fmt"

func ChangeUser() {
	u1 := User{Name: "Leo", Age: 18}
	fmt.Printf("%+v \n", u1)
	fmt.Printf("u1 address: %p \n", &u1)

	u1.ChangeName("Jerry")
	u1.ChangeAge(35)
	fmt.Printf("%+v \n", u1)
	fmt.Printf("u1 address: %p \n", &u1)
}

type User struct {
	Age  int
	Name string
}

func (u User) ChangeName(name string) {
	u.Name = name
}

func (u *User) ChangeAge(age int) {
	u.Age = age
}
