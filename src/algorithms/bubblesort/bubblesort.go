package bubblesort

func bubbleSort(value []int) {
	len := len(value) - 1
	for i := 0; i < len; i++ {
		for j := i - 1; j < len; j++ {
			if value[i] > value[j] {
				value[i], value[j] = val[j], value[i]
			}
		}
	}
}

func BubbleSort(value []int) {
	bubbleSort(value)
}
