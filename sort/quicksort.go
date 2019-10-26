package sort

func QuickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	var pivot int
	pivot = findPivot(a)

	lastIndex := len(a) - 1
	for i := 0; i <= lastIndex; i++ {
		if i == lastIndex {
			pivot = findPivot(a)
		}
		if a[i] >= pivot {
			for i <= lastIndex {
				if a[lastIndex] < pivot {
					tmp := a[i]
					a[i] = a[lastIndex]
					a[lastIndex] = tmp
					break
				}
				lastIndex--
			}
		}
	}
	return a
}

func findPivot(a []int) int {
	if a[0] > a[1] {
		return a[0]
	} else {
		return a[1]
	}
}
