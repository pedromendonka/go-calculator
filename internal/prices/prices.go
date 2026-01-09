// Package prices provides utilities for handling pricing conversion.
package prices

import (
	"fmt"
	"math"

	"go-calculator/internal/conversion"
	"go-calculator/internal/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]float64  `json:"tax_included_prices"`
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: iom,
		TaxRate:   taxRate,
	}
}

func (job *TaxIncludedPriceJob) Calculate() {
	job.LoadData()

	results := make(map[string]float64)

	for _, price := range job.InputPrices {
		value := price * (1 + job.TaxRate)
		rounded := math.Round(value*100) / 100
		results[fmt.Sprintf("%.2f", price)] = rounded
	}

	job.TaxIncludedPrices = results
	err := job.IOManager.WriteJSON(results)
	if err != nil {
		fmt.Println("Save JSON Error: ", err)
		return
	}
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println("Load Data Error: ", err)
		return
	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println("Conversion Error: ", err)
		return
	}

	job.InputPrices = prices
}
