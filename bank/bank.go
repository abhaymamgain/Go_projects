package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64
	fmt.Println(reflect.TypeOf(x) == reflect.TypeOf(float64(0)))
	// balance := 10000.0
	// fmt.Println("welcome to Go bank")

	// fmt.Println("1. Deposit Money")
	// fmt.Println("2. Withdraw Money")
	// fmt.Println("3. Check Balance")
	// fmt.Println("4. Exit")

	// var choice int
	// fmt.Print("Enter your choice: ")
	// fmt.Scan(&choice)
	// fmt.Println("your choice : ", choice)

	// if choice == 3 {
	// 	fmt.Printf("your balance is : %.2f", balance)
	// } else if choice == 1 {
	// 	var depositeAmount float64
	// 	fmt.Print("enter deposit amount : ")
	// 	fmt.Scan(&depositeAmount)
	// 	balance += depositeAmount
	// 	fmt.Printf("your new balance is : %.2f", balance)
	// } else if choice == 2 {
	// 	var withdrawAmount float64
	// 	fmt.Print("enter withdraw amount : ")
	// 	fmt.Scan(&withdrawAmount)
	// 	if withdrawAmount > balance {
	// 		fmt.Println("Insufficient balance")
	// 		return
	// 	} else {
	// 		balance -= withdrawAmount
	// 		fmt.Printf("your new balance is : %.2f", balance)
	// 	}

	// } else {
	// 	fmt.Print("goodbye")
	// }
}
