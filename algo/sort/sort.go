package sort

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

func Quick(arr []int) []int {

	arrSize := len(arr)
	if arrSize < 2 || arrSize == 0 {
		return arr
	}

	pivot := (arrSize / 2)

	arrBefore := []int{}
	arrAfter := []int{}
	newArray := []int{}

	for i := 0; i < arrSize; i++ {
		if arr[i] < arr[pivot] {
			arrBefore = append(arrBefore, arr[i])
		} else if arr[i] > arr[pivot] {
			arrAfter = append(arrAfter, arr[i])
		}
	}

	if len(arrBefore) != 0 {
		newArray = append(newArray, Quick(arrBefore)...)
	}

	newArray = append(newArray, arr[pivot])

	if len(arrAfter) != 0 {
		newArray = append(newArray, Quick(arrAfter)...)
	}

	return newArray
}

// recusion(x int): Inform a int number
func Recusion(x int) int {
	if x <= 1 {
		return x
	} else {
		return Recusion(x-1) * 2
	}
}
