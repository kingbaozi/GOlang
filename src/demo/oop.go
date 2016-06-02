package demo

import "fmt"

type Integer int

func Less(a, b int) {
	tmp, tmp2 := Integer(a), Integer(b)
	val := tmp.byless(tmp2)
	fmt.Println(val)
}
func (a Integer) byless(b Integer) bool {
	return a > b
}
