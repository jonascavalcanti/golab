package main

import "testing"

func SelectionSortTest(t *testing.T) {
	numb := [5]int{1, 2, 3, 4, 5}
	result := SelectionSort(numb)

	if !result {
		t.Errorf("result %d", result)
	}
}
