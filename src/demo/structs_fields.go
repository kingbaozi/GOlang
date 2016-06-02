package demo

import (
	"fmt"
	"strings"
)

type person struct {
	firstname string
	lastname  string
}
type T struct {
	id   int
	name string
}
type Intevar struct {
	start int
	end   int
}
type Node struct {
	data string
	su   *Node
}

func upPerson(p *person) {
	p.firstname = strings.ToUpper(p.firstname)
	p.lastname = strings.ToUpper(p.lastname)
}
func FieldsOut() {
	var ba T
	ba = T{123, "asdf"}
	m := new(T)
	m.id = 1
	m.name = "king"
	ms := &T{1, "bab"}
	var va person
	va.firstname = "sadf"
	va.lastname = "sdf"
	upPerson(&va)
	a := new(Node)
	a.data = "sadfsdaf"
	a.su = new(Node)
	a.su.data = "sadfasdf"
	a.su.su = new(Node)
	//fmt.Printf("The int is %d\n", m.id)
	//fmt.Printf("The int is %s\n", m.name)
	fmt.Println(va)
	fmt.Println(m)
	fmt.Println(ba)
	fmt.Println(a)
	fmt.Println(a.su)

	fmt.Println(ms)
}
