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
	a := []int{-317, -381, -14, -215, -590, -243, -412, 380, -312, 925, 158, -46, 177, 22, -482, 273, 217, 514, -392, 424}
	sorted := QuickSort(a)
	println(fmt.Sprintf("complete: a=%v", a))
	expected := []int{-590, -482, -412, -392, -381, -317, -312, -243, -215, -46, -14, 22, 158, 177, 217, 273, 380, 424, 514, 925}
	if !reflect.DeepEqual(sorted, expected) {
		t.Error(fmt.Sprintf("Not correct sorted: sorted=%v, expected=%v", sorted, expected))
	}
}
