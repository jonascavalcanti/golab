package algthms

import "fmt"

//SelectionSort(arr []int): inform a int array
func SelectionSort(arr []int) {

	fmt.Println("BIG-O: O(n2)")

	fmt.Println("Initial Array:", arr)
	var n = len(arr)

	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}

	fmt.Println("Array ordened:", arr)
}

//recusion(x int): Inform a int number
func Recusion(x int) int {
	if x <= 1 {
		return x
	} else {
		return Recusion(x-1) + Recusion(x-2)
	}
}
