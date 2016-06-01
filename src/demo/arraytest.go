package demo

import (
	"fmt"
)

func SumCallback() {
	var sumarray = [3]float64{1.2, 3.2, 4.4}
	allsum := sum(&sumarray)
	fmt.Println(allsum)
}

func sum(a *[3]float64) (sum float64) {
	for _, val := range a {
		sum += val
	}
	return
}

func SliceTest() {
	x := []int{1, 2, 3, 4, 5}
	y := x[0:1]
	for _, val := range y {
		fmt.Println(val)
	}
}

func ReSlicing() {
	slice1 := make([]int, 0, 10)
	for i := 0; i < cap(slice1); i++ {
		slice1 = slice1[0 : i+1]
		slice1[i] = i
		fmt.Printf("The length of slice is %d\n", len(slice1))
	}
	for _, val := range slice1 {
		fmt.Printf("The length of slice is %d\n", val)

	}
}

func CopyAppend() {
	sl_frm := []int{1, 2, 3}
	sl_to := make([]int, 10)
	n := copy(sl_to, sl_frm)
	fmt.Println(sl_to)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
}
func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { //总长度大于容量
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}
func Filters() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	bk_filter(arr, fn)
}
func bk_filter(s []int, fn func(int) bool) {
	for _, val := range s {
		if fn(val) {
			fmt.Println(val)
		}
	}
}

func fn(c int) (a bool) {
	if c%2 != 0 {
		a = false
	} else {
		a = true
	}
	return
}

func StrArr() {
	s := "hello"
	c := []byte(s)
	c[0] = 'c'
	s = string(c)
	fmt.Println(s)
}

func SliceStingbak(str string, i int) (filst_str string, last_str string) {
	if i > len(str) {
		return
	}
	s := []byte(str)
	filst_str, last_str = string(s[0:i-1]), string(s[i-1:])
	return
}
