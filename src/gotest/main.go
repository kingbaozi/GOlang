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
	val := &demo.Kc{demo.Point{3,4}, "pykk"}

	fmt.Println(val.Abs());
	//longCalculation()
	end := time.Now()
	delta := end.Sub(start)
	fmt.Println(delta)
}
