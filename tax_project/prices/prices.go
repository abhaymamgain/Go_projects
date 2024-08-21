package prices

import (
	"example/tax/conversions"
	"example/tax/filemanager"

	"fmt"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64           `json:"tax_rate"`
	Prices            []float64         `json:"prices"`
	TaxIncludedPrices map[string]string `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := filemanager.ReadLine()

	job.Prices, err = conversions.StringsToFloat(lines)
	if err != nil {
		fmt.Println("error parsing float:", err)

	}
}
func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)
	for _, price := range job.Prices {
		taxIncluded := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncluded)
	}
	job.TaxIncludedPrices = result
	filemanager.WriteFile(job, fmt.Sprintf("results_%v.json", job.TaxRate))
	fmt.Println(job)
}
func NewTaxIncludedPriceJob(tax float64) *TaxIncludedPriceJob { //contructor
	prices := []float64{}
	return &TaxIncludedPriceJob{
		Prices:            prices,
		TaxRate:           tax,
		TaxIncludedPrices: make(map[string]string),
	}
}
