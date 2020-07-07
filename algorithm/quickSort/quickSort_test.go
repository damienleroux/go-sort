package gosort

import (
	"go-sort/mock"
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
	testSamples := mock.GetTestSamples()

	for _, test := range testSamples {
		testEntry(t, test.Input, test.Expected)
	}
}
