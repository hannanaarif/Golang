package main

import "fmt"

func main(){
	websites:=map[string]string{
		"Google":"www.google.com.",
		"Amazon webservice":"www.aws.com",
	}

	fmt.Println(websites["Google"])
	fmt.Println(websites["Amazon webservice"])
	websites["LinkedIn"]="https://www.linkedIn.com"
	fmt.Println(websites)
	delete(websites,"Google")
	fmt.Println(websites)
}