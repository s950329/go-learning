package main

func Map() {
	m1 := map[string]string{
		"key1": "value1",
	}
	println(m1)

	m2 := make(map[string]string, 4)
	//m2["key2"] = "value2"
	println(m2)

	val1, ok := m1["key1"]
	println(val1, ok)
	val2, ok := m2["key2"]
	println(val2, ok)

	val2 = m2["key2"]
	println(val2)

	val2 = m2["key1"]
	println(val2)

	println("len m1: ", len(m1))
	println("len m2: ", len(m2))

	for k, v := range m1 {
		println(k, v)
	}

	delete(m1, "key1")

}