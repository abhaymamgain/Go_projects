package main

import (
	"example/tax/prices"
)

func main() {

	taxRate := []float64{.2, .3, .5}

	for _, tax := range taxRate {
		job := prices.NewTaxIncludedPriceJob(tax)
		job.Process()
	}

}
