package sort

import "fmt"

func MergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	println(fmt.Sprintf("array: %v", a))

	center := len(a) / 2
	leftSorted := MergeSort(a[:center])
	tempLeft := make([]int, len(leftSorted))
	for i := 0; i < len(leftSorted); i++ {
		tempLeft[i] = leftSorted[i]
	}

	rightSorted := MergeSort(a[center:])
	tempRight := make([]int, len(rightSorted))
	for i := 0; i < len(rightSorted); i++ {
		tempRight[i] = rightSorted[i]
	}

	i, left, right := 0, 0, 0
	println(fmt.Sprintf("Sort tempLeft=%v and tempRight%v", tempLeft, tempRight))
	for left < len(tempLeft) || right < len(tempRight) {
		if left >= len(tempLeft) && right < len(tempRight) {
			a[i] = tempRight[right]
			right++
		} else if right >= len(tempRight) && left < len(tempLeft) {
			a[i] = tempLeft[left]
			left++
		} else if tempLeft[left] < tempRight[right] {
			a[i] = tempLeft[left]
			left++
		} else if tempLeft[left] > tempRight[right] {
			a[i] = tempRight[right]
			right++
		}
		i++
	}
	println(fmt.Sprintf("Sorted a=%v", a))

	return a
}
