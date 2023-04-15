package ch1

import (
	"fmt"
	"os"
	"strings"
)

func Echo(args []string) {
	var s string

	//First format
	for _, args := range os.Args[1:] {
		s += args + " "
	}
	fmt.Println(s)

	//Second format
	fmt.Println(strings.Join(os.Args[1:], " "))

	fmt.Println("Exercise 1.1:")
	if len(os.Args) < 2 {
		fmt.Printf("Example: %s <string 1> <string 2> ...", os.Args[0])
		return
	}

	fmt.Println("Exercise 1.2:")
	for idx, args := range os.Args[1:] {
		fmt.Println(idx, ":", args)

	}

}
