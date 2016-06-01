package demo

import (
	"fmt"
<<<<<<< HEAD
	//"strings"
=======
	"strings"
>>>>>>> 037eb39d41c2b1f9e8fc69996100cb15b4a94354
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
