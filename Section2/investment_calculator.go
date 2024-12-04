package main

import (
	"fmt"
	"math"
)

func main() {
    inflationrate := 2.5
	investmentAmount := 1000.0
	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount);
	expectedReturnRate := 5.5
	year := 10.00

	fmt.Print("Enter Years: ");
	fmt.Scan(&year);
	fmt.Println("Entered Year: ",year)
	var futureValue = investmentAmount* math.Pow(1+expectedReturnRate/100,year);
	fmt.Println("Inflation ",futureValue*inflationrate);
	fmt.Println(futureValue);
}