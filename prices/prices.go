package prices

import (
	"com.github.rakhmedovrs/go-practice-project/conversion"
	"com.github.rakhmedovrs/go-practice-project/iomanager"
	"fmt"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
	IOManager         iomanager.IOManager `json:"-"`
}

func NewTaxIncludedPriceJob(manager iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:           taxRate,
		TaxIncludedPrices: make(map[string]string),
		IOManager:         manager,
	}
}

func (job *TaxIncludedPriceJob) Process() error {
	err := job.LoadData()
	if err != nil {
		return err
	}
	pricesWithTaxes := make(map[string]string)
	for _, price := range job.InputPrices {
		priceWithTaxes := price * (1 + job.TaxRate)
		pricesWithTaxes[strconv.FormatFloat(price, 'f', -1, 64)] = fmt.Sprintf("%.2f", priceWithTaxes)
	}
	job.TaxIncludedPrices = pricesWithTaxes
	err = job.IOManager.WriteDataAsJson(job)
	if err != nil {
		return err
	}
	return nil
}

func (job *TaxIncludedPriceJob) LoadData() error {
	stringPrices, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}
	floatPrices, err := conversion.StringsToFloats(stringPrices)
	if err != nil {
		return err
	}
	job.InputPrices = floatPrices
	return nil
}
