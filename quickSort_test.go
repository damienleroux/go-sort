package main

import (
	"reflect"
	"testing"
)

func testEntry(t *testing.T, entry []int, output []int) {
	QuickSort(&entry)
	if !reflect.DeepEqual(entry, output) {
		t.Errorf("%v is not equal to %v", entry, output)
	}
}

func TestQuickSort(t *testing.T) {
	testEntry(t, []int{1}, []int{1})
	testEntry(t, []int{1, 2}, []int{1, 2})
	testEntry(t, []int{2, 1}, []int{1, 2})
	testEntry(t, []int{3, 2, 1}, []int{1, 2, 3})
	testEntry(t, []int{2, 3, 1}, []int{1, 2, 3})
	testEntry(t, []int{2, 1, 3}, []int{1, 2, 3})
	testEntry(t, []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5})
	testEntry(t, []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5})
	testEntry(t, []int{1, 3, 2, 4, 5}, []int{1, 2, 3, 4, 5})
	testEntry(t, []int{5, 3, 1, 2, 4}, []int{1, 2, 3, 4, 5})
	testEntry(t, []int{}, []int{})
}
