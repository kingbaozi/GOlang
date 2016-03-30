package qsort

func quickSotr(values []int, left, right int) {
	temp := values[left]
	p := left
	i, j := left, right
	for i <= j {
		for j >= p && values[j] >= tmp {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}
		if values[i] <= temp && i <= p {
			i++
		}
		if i <= p {
			values[p] = values[i]
			p = i
		}
		values[p] = temp
		if p-left > 1 {
			quickSotr(values, left, p-1)
		}
		if right-p > 1 {
			quickSotr(values, p+1, right)
		}
	}
}
func QuickSotr(values []int) {
	quickSotr(values, 0, len(values)-1)
}
