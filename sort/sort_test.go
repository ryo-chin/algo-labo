package sort

import (
	"fmt"
	"reflect"
	"testing"
)

var a = []int{-317, -381, -14, -215, -590, -243, -412, 380, -312, 925, 158, -46, 177, 22, -482, 273, 217, 514, -392, 424}
var expected = []int{-590, -482, -412, -392, -381, -317, -312, -243, -215, -46, -14, 22, 158, 177, 217, 273, 380, 424, 514, 925}

func TestBubbleSort(t *testing.T) {
	sorted := BubbleSort(a)
	if !reflect.DeepEqual(sorted, expected) {
		t.Error(fmt.Sprintf("Not correct sorted: sorted=%v, expected=%v", sorted, expected))
	}
}

func TestQuickSort(t *testing.T) {
	sorted := QuickSort(a)
	println(fmt.Sprintf("complete: a=%v", a))
	if !reflect.DeepEqual(sorted, expected) {
		t.Error(fmt.Sprintf("Not correct sorted: sorted=%v, expected=%v", sorted, expected))
	}
}

func TestMergeSort(t *testing.T) {
	sorted := MergeSort(a)
	println(fmt.Sprintf("complete: a=%v", a))
	if !reflect.DeepEqual(sorted, expected) {
		t.Error(fmt.Sprintf("Not correct sorted: sorted=%v, expected=%v", sorted, expected))
	}
}
