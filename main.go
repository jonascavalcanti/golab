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
	fmt.Println("BIG-O: O(n2)")
	fmt.Println("Initial Array:", arr)
	fmt.Println("Array ordened: ", alg.SelectionSort(arr))

	fmt.Println("\n")

	arr = []int{200, 5, 50, 1, 2, 20, 100}
	fmt.Println("Quick Sort")
	fmt.Println("BIG-O: O(n.logn)")
	fmt.Println("Initial Array:", arr)
	fmt.Println("Array ordened: ", alg.QuickSort(arr))

	fmt.Println("\n")

	fmt.Println("Recursion (n-1) * 2")
	fmt.Println("BIG-O: O(n)")
	fmt.Println(alg.Recusion(7))

	fmt.Println("\n")

	fmt.Println("Section: Manipulator Types")
	fmt.Println("String Manipulate")
	mt.StringM("test")

}
