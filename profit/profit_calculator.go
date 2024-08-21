package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	revenue, err := takeInput("revenue")
	if err != nil {
		fmt.Println(err)
		return
	}
	expenses, err := takeInput("expenses")
	if err != nil {
		fmt.Println(err)
		return
	}
	taxRate, err := takeInput("tax rate")
	if err != nil {
		fmt.Println(err)
		return
	}
	earningBeforeTax, earningsAfterTax, ratio := calculatevalues(revenue, expenses, taxRate)
	fmt.Printf("earning before tax %.2f\n", earningBeforeTax)
	fmt.Printf("earning after tax %.2f\n", earningsAfterTax)
	fmt.Printf("ratio %.2f\n", ratio)
}

func calculatevalues(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	pv := revenue - expenses
	taxAmount := pv * (taxRate / 100)
	rpv := pv - taxAmount
	ratio := pv / rpv
	return pv, rpv, ratio

}

func takeInput(input string) (float64, error) {

	var tempinput string
	fmt.Printf("%v : ", input)
	fmt.Scan(&tempinput)
	fmt.Print(tempinput)
	tempinput1, err := strconv.ParseFloat(tempinput, 64)
	fmt.Println(reflect.TypeOf(tempinput1))
	fmt.Println(reflect.TypeOf(float64(0)))

	if err != nil {
		fmt.Print(err.Error())
		return 0, errors.New("\nentered value must be a number")
	}
	return tempinput1, nil
}
