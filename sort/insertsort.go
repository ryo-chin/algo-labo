package sort

func InsertSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		j := i + 1
		for j > 0 && j < len(a) && a[j-1] > a[j] {
			a[j-1], a[j] = a[j], a[j-1]
			j--
		}
	}
	return a
}
