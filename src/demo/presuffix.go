package demo

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var str string = "hello world!!"

func Presuffix() {
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str, "Th")
	//是否已Th 开头
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
	//是否已!! 结尾
	fmt.Printf("%t\n", strings.HasSuffix(str, "!!"))
}

//判断子字符串或字符在父字符串中出现的位置（索引）
func Index_in_string() {
	fmt.Printf("The position of \"or\" is: ")
	//or 第一次出现的位置
	fmt.Printf("%d\n", strings.Index(str, "or"))

	//o 第一次出现的位置
	fmt.Printf("The position of the first instance of \"o\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "o"))
	//o 最后一次出现的位置
	fmt.Printf("The position of the last instance of \"o\" is: ")
	fmt.Printf("%d\n", strings.LastIndex(str, "o"))
	//不包含是返回 -1
	fmt.Printf("The position of \"Burger\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Burger"))
	//字符串替换
	fmt.Printf("The Replace of \"l\" ,\"i\" is: ")
	str = strings.Replace(str, "l", "i", -1)
	fmt.Println(str)
}
func CountSubstring() {
	var manyG = "gggggggggg"
	fmt.Printf("%d\n", strings.Count(str, "l"))
	fmt.Printf("%d\n", strings.Count(manyG, "g"))
}

func SliceSting() {
	str := "The quick brown fox jumps over the lazy dog"
	//利用 1 个或多个空白符号来作为动态长度的分隔符将字符串分割成若干小块，并返回一个 slice
	sl := strings.Fields(str)
	fmt.Printf("Splitted in slice: %v\n", sl)
	for _, val := range sl {
		fmt.Println(val)
	}
	fmt.Println()
	str2 := "GO1|The ABC of Go|25"
	//strings.Split(s, sep) 用于自定义分割符号来对指定字符串进行分割，同样返回 slice。
	sl2 := strings.Split(str2, "|")
	fmt.Printf("Splitted in slice: %v\n", sl)
	for _, val := range sl2 {
		fmt.Println(val)
	}
	//Join 用于将元素类型为 string 的 slice 使用分割符号来拼接组成一个字符串：
	fmt.Println()
	str3 := strings.Join(sl, ";")
	fmt.Println(str3)
}

func StringConversion() {
	var or string = "666"
	var an int
	var nesS string
	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	//将字符串转换为 int 型
	an, _ = strconv.Atoi(or)
	fmt.Printf("The integer is : %d\n", an)
	an = an + 5
	nesS = strconv.Itoa(an)
	fmt.Printf("the new string is : %s\n", nesS)
}

func StringConversion2() {
	var (
		or   string = "ABC"
		newS string
	)
	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	an, err := strconv.Atoi(or)
	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error\n", or)
		os.Exit(1)
	}
	fmt.Printf("The integer is : %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("the new string is : %s\n", newS)
}

func Testfat() {
	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		//fallthrough
	case 5:
		fmt.Println("was <= 5")
		//fallthrough
	case 6:
		fmt.Println("was <= 6")
		//fallthrough
	case 7:
		fmt.Println("was <= 7")
		//fallthrough
	case 8:
		fmt.Println("was <= 8")
		//fallthrough
	default:
		fmt.Println("default case")
	}
}
