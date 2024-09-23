package main

import (
	"com.github.rakhmedovrs/go-practice-project/filemanager"
	"com.github.rakhmedovrs/go-practice-project/prices"
	"fmt"
)

const PricesFileName = "resources/prices.txt"
const TaxIncludedPriceJobFileName = "result_tax_rate_%.0f.json"

func main() {
	var taxRates = []float64{0, 0.03, 0.10, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New(PricesFileName, fmt.Sprintf(TaxIncludedPriceJobFileName, taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Println(err)
		}
	}
}
