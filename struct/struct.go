package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {

	var Name string = takeInput("Name")
	var Age int = takeInputInt("Age")
	var appUser user.User = user.CreateUser(Name, Age)
	var appPointer *user.User = &appUser
	appPointer.PrintUser()
	var appAdmin user.Admin = user.NewAdmin("admin", "admin", Name, Age)
	var appPointerAdmin *user.Admin = &appAdmin
	appPointerAdmin.User.PrintUser()

}

func takeInput(prompt string) string {
	var input string
	fmt.Print(prompt + ": ")
	fmt.Scanln(&input)
	return input
}

func takeInputInt(prompt string) int {
	var input int
	fmt.Print(prompt + ": ")
	fmt.Scanln(&input)
	return input
}
