package sort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	a := []int{2, 8, 6, 10, 4, 2}
	sorted := BubbleSort(a)
	expected := []int{2, 2, 4, 6, 8, 10}
	if !reflect.DeepEqual(sorted, expected) {
		t.Error(fmt.Sprintf("Not correct sorted: sorted=%v, expected=%v", sorted, expected))
	}
}

func TestQuickSort(t *testing.T) {
	a := []int{2, 8, 6, 4, 9, 10, 3, 7, -1}
	sorted := QuickSort(a)
	expected := []int{-1, 2, 3, 4, 6, 7, 8, 9, 10}
	if !reflect.DeepEqual(sorted, expected) {
		t.Error(fmt.Sprintf("Not correct sorted: sorted=%v, expected=%v", sorted, expected))
	}
}
