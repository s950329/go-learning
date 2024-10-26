package main

func Switch(age int) string {
	switch {
	case age >= 18:
		return "成年了"
	case age < 6:
		return "嬰幼兒"
	default:
		return "青少年"
	}
}
