package main

import (
	"fmt"
	"os"

	"github.com/jonascavalcantineto/golabs/gopl/ch1/handler"
)

func main() {

	fmt.Println("Program Name:", os.Args[0])

	fmt.Println()

	fmt.Println("Echo version 1 is using normal for() statement:")
	handler.Echo1(os.Args)

	fmt.Println()

	fmt.Println("Echo version 2 is using range for() statement:")
	handler.Echo2(os.Args)

	fmt.Println()

	fmt.Println("Echo version 3 is using package strings.Join() statement:")
	handlers.Echo3(os.Args)

}
