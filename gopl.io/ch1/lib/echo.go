package lib

import (
	"fmt"
	"strings"
	"time"
)

func Echo1(args []string) {
	start := time.Now()

	var s, sep string

	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Println(s)

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}

func Echo2(args []string) {
	start := time.Now()
	var s, sep string

	for idx, arg := range args[1:] {
		s += sep + arg
		sep = " "

		fmt.Println("idx: ", idx, "arg: ", arg)
	}

	fmt.Println(s)

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func Echo3(args []string) {
	start := time.Now()

	fmt.Println(strings.Join(args[1:], " "))

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}
