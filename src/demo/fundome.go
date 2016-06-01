package demo

import (
	"fmt"
)

func Args(prox int, who ...string) {
	fmt.Println(prox)
	for _, val := range who {
		fmt.Println(val)
	}
}
func ArgFunc() {
	startfunc(1, 2, add)
}

func startfunc(a, b int, f func(int, int)) {
	f(a, b)
}

func add(a, b int) {
	fmt.Printf("a + b = %d", a+b)
}
