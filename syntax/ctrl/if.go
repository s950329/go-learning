package main

func IfOnly(age int) string {
	if age >= 18 {
		return "成年了"
	}

	return "未成年"
}

func IfElse(age int) string {
	if age >= 18 {
		return "成年了"
	} else {
		return "未成年"
	}
}

func IfElseIf(age int) string {
	if age >= 18 {
		return "成年了"
	} else if age >= 12 {
		return "少年"
	} else {
		return "未成年"
	}
}

func IfNewVariable(start int, end int) string {
	if distance := end - start; distance > 100 {
		return "too far"
	} else {
		return "OK"
	}
}
