package main

import (
	"fmt"

	sorter "github.com/jonascavalcanti/golab/algo/sorter"
)

func main() {
	numbers := []int{1, 9, 5, 10}

	fmt.Println(sorter.Selection(numbers))

}
