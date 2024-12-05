package main

import "fmt"

func main(){
	age:=32
	fmt.Println("Address of age in adult func",&age);

	var agePointer *int

	agePointer =&age

	fmt.Println("Value pointed by agePointer:", *agePointer)

	adultYears :=getAdultYears(agePointer)
	fmt.Println(adultYears);


}

func getAdultYears(age *int) int {
	fmt.Println("Address of age in adult func",&age);
	return *age -18
}