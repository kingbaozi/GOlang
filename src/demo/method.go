package demo

import (
	"fmt"
)

type pron struct {
	a int
	b int
}
type IntVectot []int
type employee struct {
	salary float32
}

func StrNewTs() {
	//flag := pron{123, 55}
	a := IntVectot{3, 2, 3, 56, 213}
	//fmt.Printf("sum %d", flag.structsum())
	b := pron{1, 2}
	b.change()
	b1 := new(pron)
	b1.a = 23
	b1.b = 12
	b1.change()

	fmt.Println(b.white())
	fmt.Println(b1.white())
	fmt.Printf("add %g", employee{12888}.giveRaise(0.2))
	fmt.Printf("add %d", a.sum())
}
func (tir *pron) structsum() int {
	return tir.a + tir.b
}
func (tir *pron) additem(param int) int {
	return tir.a + tir.b + param
}
func (cess IntVectot) sum() (s int) {
	for _, n := range cess {
		s += n
	}
	return
}
func (p employee) giveRaise(per float32) float32 {
	return (p.salary + p.salary*per)
}

func (p *pron) change() {
	p.a = 2
}
func (p pron) white() string {
	return fmt.Sprint(p)
}
