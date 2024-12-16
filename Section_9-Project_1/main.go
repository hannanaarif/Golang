package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main(){
	// var prices []float64=[]float64{10,20,30}
	// var taxRates []float64=[]float64{1.9,10.2,8.0}

	//prices:=[]float64{10,20,30}
	taxRates:=[]float64{5.9,10,8.0}

	for _,taxRate := range taxRates {
		fm:=filemanager.New("prices.txt",fmt.Sprintf("result_%.0f.json",taxRate*100))
		priceJob:=prices.NewTaxIncludedPriceJob(fm,taxRate)
		priceJob.Process()
	}

	// result:=make(map[float64][]float64);


	// fmt.Println(result)
}