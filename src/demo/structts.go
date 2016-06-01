package demo

import (
	"fmt"
	//"strings"
)

type T struct {
	id   int
	name string
}

func NewT(id int, name string) *T {
	if id < 0 {
		return nil
	}
	return &T{id, name}
}
func StrNewT() {
	ms := NewT(21, "abc")
	fmt.Println(ms)
}
