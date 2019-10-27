package sort

import "fmt"

func QuickSort(a []int) []int {
	println(fmt.Sprintf("array %v", a))
	if len(a) < 2 {
		return a
	}
	pivot := findPivot(a)

	left, right := 0, len(a)-1
	leftCur, rightCur := 0, len(a)-1

	for leftCur < rightCur {
		for a[leftCur] < pivot {
			leftCur++
		}
		for a[rightCur] > pivot {
			rightCur--
		}
		if leftCur < rightCur {
			a[leftCur], a[rightCur] = a[rightCur], a[leftCur]
		}
	}
	println(fmt.Sprintf("after sort: pivot=%v, array=%v leftCur=%v, rightCur=%v", pivot, a, leftCur, rightCur))

	if left < rightCur+1 {
		QuickSort(a[:leftCur])
	}
	if leftCur < right-1 {
		QuickSort(a[leftCur+1:])
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
