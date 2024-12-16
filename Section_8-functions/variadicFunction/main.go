package main

import (
	"fmt"
)


func main(){

	sum:=sumUp(1,2,3)
	fmt.Println("hello from variadic",sum)
}

func sumUp(numbers ...int) int{
   sum:=0

   for _, num := range numbers {
	sum += num
}
return sum
}