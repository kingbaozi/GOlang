package demo

import (
	"fmt"
	"time"
)

func TimeEcho() {
	now := time.Now()
	fmt.Println(time.Now())
	fmt.Printf("%4d.%02d.%02d", now.Year(), now.Month(), now.Day())
}
