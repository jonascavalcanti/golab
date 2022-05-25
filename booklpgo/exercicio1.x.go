package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	//fmt.Println("Name:", os.Args[0])

	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("%f\n", time.Since(start).Seconds())

	start = time.Now()
	for idx, v := range os.Args[1:] {
		fmt.Println("Idx:", idx, "Value:", v)
	}
	fmt.Printf("%f\n", time.Since(start).Seconds())
}
