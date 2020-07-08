package main

import "fmt"

func main() {
	name := "jj"
	version := "1.0"

	fmt.Println("Hello Mr.:", name)
	fmt.Println("Version: ", version)

	fmt.Println("Choice your Option:")
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair")

	var command int

	fmt.Scan(&command)

	fmt.Println("O valor da variável comando é:", command)
}
