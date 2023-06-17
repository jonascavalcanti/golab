package sorter

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
