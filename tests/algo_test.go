package tests

import "testing"
import "golabs/algo"

func SelectionSortTest(t *testing.T) {
	numb := []int{1, 2, 3, 4, 5}
	result := algo.SelectionSort(numb)

	if result == nil {
		t.Errorf("result %d", result)
	}
}
