// Package calculator provides tax calculator functionalities.
package calculator

import (
	"fmt"
	"go-calculator/internal/cmdmanager"
	"go-calculator/internal/filemanager"
	"go-calculator/internal/iomanager"
	"go-calculator/internal/prices"
)

const DataDirPath = "internal/shared/data"
const PricesFilePathName = DataDirPath + "/prices.txt"
const JSONDirPath = DataDirPath + "/json/"
const JSONFilePathName = JSONDirPath + "tax_included_prices_"

func Run() {
	taxRates := []float64{0, 0.1, 0.15, 0.2}

	var option int
	fmt.Println("1. File Manager")
	fmt.Println("2. Command Manager")
	_, err := fmt.Scanln(&option)
	if err != nil {
		fmt.Println("Error reading option:", err)
		return
	}

	switch option {
	case 1:
		fmt.Println("File Manager selected")
	case 2:
		fmt.Println("Command Manager selected")
	default:
		fmt.Println("Invalid option selected")
		return
	}

	for i, taxRate := range taxRates {
		fmt.Printf("Starting job %v of %v for tax rate: %.2f\n", i+1, len(taxRates), taxRate)
		var iom iomanager.IOManager

		switch option {
		case 1:
			outputFileName := JSONFilePathName + fmt.Sprint(taxRate)
			iom = filemanager.New(PricesFilePathName, outputFileName)
		case 2:
			iom = cmdmanager.New()
		}

		pricesJob := prices.NewTaxIncludedPriceJob(iom, taxRate)
		pricesJob.Calculate()

		fmt.Printf("Job completed...\n\n")
	}
}
