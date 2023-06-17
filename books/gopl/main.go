package main

import (
	"fmt"
	"os"

	ch1 "github.com/jonascavalcanti/golab/books/gopl/ch1"
)

func main() {

	fmt.Println("Program Name:", os.Args[0])

	fmt.Println()

	fmt.Println("Echo version 1 is using normal for() statement:")
	ch1.Echo1(os.Args)

	fmt.Println()

	fmt.Println("Echo version 2 is using range for() statement:")
	ch1.Echo2(os.Args)

	fmt.Println()

	fmt.Println("Echo version 3 is using package strings.Join() statement:")
	ch1.Echo3(os.Args)

}
