package main

func Loop() {
	for i := 0; i < 10; i++ {
	}
}

func Loop2() {
	i := 0
	for i < 100 {
		i++
	}
}

func Loop3() {
	for {
		println("hello world")
	}
}

func ForArray() {
	println("遍歷數組")
	arr := [3]string{"A", "B", "C"}
	for i, v := range arr {
		println(i, v)
	}
}

func ForMap() {
	println("遍歷 map")
	m := map[string]string{
		"1": "A",
		"2": "B",
		"3": "C",
	}
	for k, v := range m {
		println(k, v)
	}
}

func LookBug() {
	users := []User{
		{
			Name: "Tom",
		},
		{
			Name: "Jack",
		},
	}

	m := make(map[string]*User)

	for _, user := range users {
		m[user.Name] = &user
	}

	for name, user := range m {
		println(name, user.Name)
	}
}

type User struct {
	Name string
}
