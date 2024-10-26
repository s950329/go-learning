package main

func Sum[T Number](values []T) T {
	var res T
	for _, v := range values {
		res = res + v
	}
	return res
}

func Max[T Number](values []T) T {
	t := values[0]
	for _, v := range values {
		if t < v {
			t = v
		}
	}
	return t
}

func Find[T any](values []T, filter func(t T) bool) T {
	for _, v := range values {
		if filter(v) {
			return v
		}
	}
	var t T
	return t
}

func Insert[T any](index int, val T, values []T) []T {
	if index < 0 || index > len(values) {
		panic("invalid index")
	}

	// extend
	values = append(values, val)
	for i := len(values) - 1; i > index; i-- {
		if i-1 > 0 {
			values[i] = values[i-1]
		}
	}
	values[index] = val

	return values
}

type Integer int

type Number interface {
	~int | uint
}

func UseSum() {
	res := Sum[int]([]int{123, 234})
	println(res)
	resV1 := Sum[Integer]([]Integer{123, 234})
	println(resV1)
}
