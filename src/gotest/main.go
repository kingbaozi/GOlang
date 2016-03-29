package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := rand.Int()
	fmt.Println(a)
	b := rand.Intn(8) //【0，n)
	timens := int64(time.Now().Nanosecond())
	//生成随机数种子
	rand.Seed(timens)
	c := 100 * rand.Float32()
}
