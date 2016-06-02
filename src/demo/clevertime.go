package demo

import (
	"fmt"
	"time"
)

func Greetings() {
	var (
		hh = time.Now().Hour()
		mm = time.Now().Minute()
		ss = time.Now().Second()
	)

	y, m, d := time.Now().Date()
	t := time.Now()
	mytime := time.Now().Format("3:05am")
	mmytime := t.Format("2006-01-02 15:04:05 pm")
	week := time.Now().Weekday().String()
	//m := time.November

	fmt.Println("Holle sir, now " + week)
	fmt.Println(mytime)
	fmt.Println(mmytime)
	fmt.Println(hh, mm, ss)
	fmt.Println(y, m, d)
}
