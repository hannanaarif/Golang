package main

import "fmt"

func main(){
	var productNames [4]string= [4]string {"A Book"} 
	prices:=[4]float64{10.99,20.22,25.65,45.12}
	productNames[2]="A carpet"
	fmt.Println(prices)
	fmt.Println(productNames)
	fmt.Println(prices[2])
	featuredPrice:=prices[1:]
	highlightedPrice:=prices[:1]
	fmt.Println(highlightedPrice)
	fmt.Println(featuredPrice)
	fmt.Println(len(featuredPrice))
}