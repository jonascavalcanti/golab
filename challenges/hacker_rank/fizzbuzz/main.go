package main

import "fmt"

func main() {
	n := int32(15)
	fizzbuzz(n)
}

func fizzbuzz(n int32) {
	for i := 1; i <= int(n); i++ {
		if (i%3 == 0) && (i%5 == 0) {
			fmt.Println("FizzBuzz")
		}
		if (i%3 == 0) && (i%5 != 0) {
			fmt.Println("Fizz")
		}
		if (i%3 != 0) && (i%5 == 0) {
			fmt.Println("Buzz")
		}
		if (i%3 != 0) && (i%5 != 0) {
			fmt.Println(i)
		}

	}
}
