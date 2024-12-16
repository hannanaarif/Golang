package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/filemanager"
)

type TaxIncludePricesJob struct {
	IOManager         filemanager.FileManger
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

func (job *TaxIncludePricesJob) LoadData() {

	lines, err := job.IOManager.Readfile()

	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringToFloats(lines)

	if err != nil {
		fmt.Println("coverting price to float failed.")
		fmt.Println(err)
		return
	}
	//   prices[listIndex]=floatPrices
	job.InputPrices = prices
	job.IOManager.WriteResult(job)
}

func (job *TaxIncludePricesJob) Process() {
	job.LoadData()
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludePrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%0.2f", price)] = fmt.Sprintf("%0.2f", taxIncludePrice)
	}
	job.TaxIncludedPrices = result
	job.IOManager.WriteResult(job)
}

func NewTaxIncludedPriceJob(fm filemanager.FileManger, taxRate float64) *TaxIncludePricesJob {
	return &TaxIncludePricesJob{
		IOManager:   fm,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
