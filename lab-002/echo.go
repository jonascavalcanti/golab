package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	var str, s string

	startTime := time.Now()

	if len(os.Args) <= 1 {
		fmt.Println("ex.: ", os.Args[0], "<some arguments>")
	}

	// fmt.Println("Using range method - return a index and a value")
	// for i, arg := range os.Args {
	// 	fmt.Println("Pos[", i, "]: ", arg)

	// }
	// fmt.Println(time.Since(startTime).Seconds(), "Time execution Range")

	fmt.Println("Using for{} method")
	for i := 1; i < len(os.Args); i++ {
		str += s + os.Args[i]
		s += " "

	}
	fmt.Println(time.Since(startTime).Seconds(), "Time execution for{} method")
	fmt.Println(str)

	// fmt.Println("Using string.Join() method")
	// fmt.Println(strings.Join(os.Args[1:], " "))
	// fmt.Println(time.Since(startTime).Seconds(), "Time execution strings.Join() method")

}
