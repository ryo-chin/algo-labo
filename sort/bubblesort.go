package sort

func BubbleSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				tmp := a[j]
				a[j] = a[i]
				a[i] = tmp
			}
		}
	}
	return a
}
