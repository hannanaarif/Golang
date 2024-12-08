package main

import "fmt"


func printintSlice(items[]int){
   for _, item:=range items {
	fmt.Println(item)
   }
}

func printstringSlice(items[]string){
	for _, item:=range items {
	 fmt.Println(item)
	}
 }

 //Generics
 func printTSlice[T any](items[]T){
	for _, item:=range items {
	 fmt.Println(item)
	}
 }

 func printSlice[T int|string|bool](items[]T){
	for _, item:=range items {
	 fmt.Println(item)
	}
 }

 type stack[T any] struct{
	element[]T
 }



func main() {
	fmt.Println("Generics")

	// myStack:=stack{
	// 	element:[]int{1,2,3},
	// }

	// fmt.Println(myStack);



	myStack:=stack[string]{
		element:[]string{"golang","typescript","C++"},
	}

	fmt.Println(myStack);



	// nums:=[]int{1,2,3}
	// names:=[]string{"golang,typescript"}
	// // printstringSlice(names)
	// printSlice(names)
	// printSlice(nums)


}