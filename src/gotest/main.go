package main

import (
	"demo"
	"fmt"
	"log"
	"runtime"
	"time"
)

var where = func() {
	_, file, line, _ := runtime.Caller(1)
	fmt.Println()
	log.Printf("%s:%d", file, line)
}
var start = time.Now()

func main() {
	val := new(demo.Person)
	fmt.Println(val.firstname)
	//longCalculation()
	end := time.Now()
	delta := end.Sub(start)
	fmt.Println(delta)
}
