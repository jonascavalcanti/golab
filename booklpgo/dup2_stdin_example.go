package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	lines := make(map[string]int)
	files := make(map[string]map[string]int)

	var count int

	for _, filename := range os.Args[1:] {

		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
			break
		}
		for _, line := range strings.Split(string(data), "\n") {
			count++
			lines[line] = count
		}
		files[filename] = lines

		fmt.Println(files)

		for line, n := range lines {
			if n >= 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}

		}
	}
}
