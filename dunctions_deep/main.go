package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	doubledNumbers := doubleNumbers(&numbers)
	fmt.Print(doubledNumbers)
}

func doubleNumbers(numbers *[]int) []int {
	dnum := []int{}
	for _, num := range *numbers {
		dnum = append(dnum, num*2)
	}
	return dnum
}
