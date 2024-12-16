package main

import (
	"fmt"
)

func main() {

	hobbies := [3]string{"coding","cooking","cuddling"}
	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	hobbies1:=hobbies[0:2]
	fmt.Println(hobbies1)

	hobbies2:=hobbies[0:2]
	fmt.Println(hobbies2)
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
	hobslice:=[]int{1,2};

	hobslice=append(hobslice,5)
	fmt.Println("hobslice: ",hobslice)

	//4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.

	hobslice1:=[]int {hobslice[1],hobslice[len(hobslice)-1]}
	fmt.Println(hobslice1)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)

	goals:=[]string {"Interview ,","morningPerson"}
	fmt.Println(goals)

	// 6) Set the second goal to a different one AND then add a third goal to 
	//that existing dynamic array
    goals[1]="Makingmovie"
	goals=append(goals, ", Biz")
	fmt.Println(goals)

	// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.

        type Product struct{
			title string
			id float64
			price float64
		}

		Products:=[]Product {{title: "Phone",id:123,price: 12345},{title: "pencil",id:13,price: 145}}

		Products=append(Products,Product{title:"Pen",id:12,price:12314})


		fmt.Println(Products);




}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.


// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.