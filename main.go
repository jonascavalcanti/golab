package main

import (
	"fmt"

	alg "github.com/jonascavalcantineto/golabs/algthms"

	mt "github.com/jonascavalcantineto/golabs/manipulator"
)

func main() {

	fmt.Println("Section: ALGORITHMS \n")

	arr := []int{5, 50, 1, 2}

	fmt.Println("Selection Sort")
	alg.SelectionSort(arr)
	fmt.Println("\n")
	fmt.Println("Recursion")
	fmt.Println(alg.Recusion(6))

	fmt.Println("Section: TYPES \n")
	fmt.Println(mt.StringM("testss"))

}
