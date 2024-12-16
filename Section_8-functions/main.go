package main

import (
	"fmt"
)

type transformFun func(int) int

func main(){
	numbers:=[]int{1,2,3,4}
	doubled:=transformNumbers(&numbers,double)
	fmt.Println(doubled)
	tripled:=transformNumbers(&numbers,triple)
	fmt.Println(tripled)
	transformer := createTransformer(2) // Create a closure with factor = 2
    fmt.Println(transformer(5))  
}

func transformNumbers(numbers *[]int, transform transformFun) []int{
	dnumbers:=[]int{}
	for _,val := range *numbers{
		dnumbers=append(dnumbers,transform(val))
	}
	return dnumbers
}

func double (number int) int {
	return number*2
}

func triple (number int) int {
	return number *3
}


func createTransformer(factor int) func(int) int {
	return func(number int) int{
		return number*factor
	}
}
