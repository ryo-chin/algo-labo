package sort

import "fmt"

func QuickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	pivot := findPivot(a)

	left, right := 0, len(a)-1
	for left < right {
		for a[left] < pivot {
			left++
		}
		for a[right] > pivot {
			right--
		}
		if left < right {
			a[left], a[right] = a[right], a[left]
		}
	}
	println(fmt.Sprintf("  sort by pivot=%v > cursor=%v, a[:cursor]=%v, a[cursor]=%v, a[cursor+1:]=%v", pivot, left, a[:left], a[left], a[left+1:]))

	QuickSort(a[:left])
	QuickSort(a[left+1:])

	return a
}

func findPivot(a []int) int {
	if a[0] > a[1] {
		return a[0]
	} else {
		return a[1]
	}
}
