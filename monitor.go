package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	showIntroduction()
	showMenu()

	// if command == 1 {
	// 	fmt.Println("Monitorando....")
	// } else if command == 2 {
	// 	fmt.Println("Exibindo logs....")
	// } else if command == 3 {
	// 	fmt.Println("Saindo...")
	// } else {
	// 	fmt.Println("Não conheço este comando")
	// }

	cmd := scanCommand()

	switch cmd {
	case 1:
		monitoring()
	case 2:
		fmt.Println("Exibindo logs....")
	case 3:
		fmt.Println("Saindo...")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
	}
}
func showIntroduction() {
	name := "JJ"
	version := "1.0"

	fmt.Println("Hello Mr.:", name)
	fmt.Println("Version: ", version)
}

func showMenu() {
	fmt.Println("Choice your Option:")
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair")
}

func scanCommand() int {
	var cmd int
	fmt.Scan(&cmd)
	return cmd
}

func monitoring() {

	fmt.Println("Monitorando....")
	site := "http://www.funceme.br"
	resp, err := http.Get(site)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
}
