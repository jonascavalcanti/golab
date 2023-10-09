package sorter

// SelectionSort(arr []int): inform a int array
func Selection(arr []int) []int {

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

	return arr
}
