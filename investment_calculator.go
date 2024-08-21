package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount, years float64 = 0, 10
	expectedReturnRate := 5.5
	fmt.Scan((&investmentAmount))
	futureValue := float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	futureRealvalue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealvalue)

}
