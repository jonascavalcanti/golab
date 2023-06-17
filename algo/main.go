package main

import (
	sorter "command-line-arguments/Users/jbenigno/Dev/Repos/Github/jonascavalcantineto/golabs/algo/sorter/recursion.go"
	"fmt"
)

func main() {
	numbers := []int{1, 9, 5, 10}

	fmt.Println(sorter.Selection(numbers))

}
