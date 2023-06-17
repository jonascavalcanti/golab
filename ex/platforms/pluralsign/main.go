package main

import (
	"fmt"
	model "github.com/jonascavalcantineto/golab/ex/platforms/pluralsign/model"
)

func main() {
	user := model.User{}

	user.ID = 1
	user.FirstName = "Jonas"
	user.LastName = "Cavalcanti"

	fmt.Println(user.FirstName)

}