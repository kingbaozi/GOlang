package demo

import (
	"fmt"
)

func MakeMap() {
	var mapLit map[string]int
	var mapAssigned map[string]int
	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3

	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])
}

func Mapfunc() {
	mf := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 10 },
		5: func() int { return 10 },
	}
	fmt.Println(mf)
}

func MapForrange() {
	items := make([]map[int]int, 5)
	for value := range items {
		items[value] = make(map[int]int, 1)
		items[value][1] = 2
	}
	fmt.Println(items)
	items2 := make([]map[int]int, 5)
	for _, value := range items2 {
		value = make(map[int]int, 1)
		value[1] = 2
	}
	fmt.Println(items2)
}

//barVal
/*
var (
    barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
                            "delta": 87, "echo": 56, "foxtrot": 12,
                            "golf": 34, "hotel": 16, "indio": 87,
                            "juliet": 65, "kili": 43, "lima": 98}
)
*/
func SortMap() {
	var (
		barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
			"delta": 87, "echo": 56, "foxtrot": 12,
			"golf": 34, "hotel": 16, "indio": 87,
			"juliet": 65, "kili": 43, "lima": 98}
		i int = 0
	)
	fmt.Println("unsorted:")
	keys := make([]map[string]int, len(barVal))
	for k, v := range barVal {
		keys[i] = make(map[string]int, 1)
		keys[i][k] = v
		i++
	}
	fmt.Println()
	fmt.Println("sorted:")
	for k, v := range keys {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}
}
