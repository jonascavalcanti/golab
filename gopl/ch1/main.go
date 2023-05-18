package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Program Name:", os.Args[0])

	fmt.Println()

	fmt.Println("Echo version 1 is using normal for() statement:")
	lib.Echo1(os.Args)

	fmt.Println()

	fmt.Println("Echo version 2 is using range for() statement:")
	lib.Echo2(os.Args)

	fmt.Println()

	fmt.Println("Echo version 3 is using package strings.Join() statement:")
	lib.Echo3(os.Args)

}
