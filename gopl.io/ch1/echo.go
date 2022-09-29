package ch1

import "fmt"

func Echo(args []string) {
	var s string

	for i := 1; i < len(args); i++ {
		s += args[i] + " "
	}

	fmt.Println(s)
}
