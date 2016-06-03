package demo

import (
	"fmt"
	//"strings"
	"reflect"
)

type TagType struct { // tags
	field1 bool   "kill 4"
	field2 string "The name of the stirng"
	field3 int    "How much there are"
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

func TagOut() {
	tt := TagType{true, "Barak Obama", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}
