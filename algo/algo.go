package algo

//{5, 50, 1, 2}
func SelectionSort(arr []int) {

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
}

func Recusion(x int) int {
	if x <= 1 {
		return x
	} else {
		return Recusion(x-1) + Recusion(x-2)
	}
}
