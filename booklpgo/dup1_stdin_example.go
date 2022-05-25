package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := make(map[string]int)
	var count int
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		if input.Text() == "" {
			break
		}
		count++
		lines[input.Text()] = count

	}

	for line, n := range lines {
		if n >= 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
