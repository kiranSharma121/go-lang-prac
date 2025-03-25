package main

import "github.com/project/price"

func main() {

	taxRates := []float64{0, 0.07, 0.1, 0.15}
	for _, taxrate := range taxRates {
		priceJob := price.NewTaxIncludedPriceJob(taxrate)
		priceJob.Process()
	}
}
