package main

import "fmt"

func main() {
	citys := make(map[int]string)

	citys[1] = "Fortaleza"
	citys[2] = "São Paulo"
	citys[3] = "Noronha"

	fmt.Println("Citys:", citys)

	valor1, isValue := citys[1]
	fmt.Println("Valor1:", valor1, isValue)

	valor2, isValue := citys[3]
	fmt.Println("Valor2:", valor2, isValue)

	delete(citys, 2)

	for i, city := range citys {
		fmt.Println("City with the number:", i, "and name", city)
	}

	beachs := map[int]string{
		1: "Noronha",
		2: "Hawaii",
		3: "Mentawaii",
	}

	fmt.Println("Beachs: ", beachs)
	fmt.Println("Beach Name: ", beachs[2])

	for i, beach := range beachs {
		fmt.Println("Beach with the number:", i, "and name", beach)
	}

	//Nil Map - this declation broken the code
	// var names map[int]string

	// names[1] = "Jonas"
	// names[2] = "Dani"

	// fmt.Println("Names: ", names)
}
