package price

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) Loadder() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("couldn't open the file")
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("reading the file contain failed")
		fmt.Println(err)
		file.Close()
		return
	}
	prices := make([]float64, len(lines))
	for lineIndex, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("converting price to float failed")
			fmt.Println(err)
			file.Close()
			return
		}
		prices[lineIndex] = floatPrice
	}
	job.InputPrices = prices

}

func (job TaxIncludedPriceJob) Process() {
	job.Loadder()
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:           taxRate,
		InputPrices:       []float64{10, 20, 30},
		TaxIncludedPrices: make(map[string]float64),
	}
}
