package main

import (
	sort "algo/sort"
	"fmt"
)

func main() {
	numbers := []int{1, 9, 5, 10}

	fmt.Println(sort.Selection(numbers))

}
