package main

type ListV1[T any] interface {
	Add(index int, val T)
	Append(val T) error
	Delete(index int) error
}

type LinkListV1[T any] struct {
	head *nodeV1[T]
}

func (l *LinkListV1[T]) Add(index int, val T) {

}

type nodeV1[T any] struct {
	data T
}

func UseList() {
	l := &LinkListV1[int]{}
	l.Add(1, 123)
	//l.Add(1, "123")
}
