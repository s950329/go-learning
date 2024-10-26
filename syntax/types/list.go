package main

type List interface {
	Add(index int, val any)
	Append(val any) error
	Delete(index int) error
}

type LinkList struct {
}

func (l *LinkList) Append(val any) error {
	//TODO implement me
	panic("implement me")
}

func (l *LinkList) Delete(index int) error {
	//TODO implement me
	panic("implement me")
}

func (l *LinkList) Add(index int, val any) {

}

type node struct {
	next *node
}
