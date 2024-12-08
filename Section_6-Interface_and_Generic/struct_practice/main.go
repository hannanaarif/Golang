package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/todo/todo"
)

func main() {

	todoText := getTodoData()

	userTodo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}
	userTodo.Display()
	userTodo.Save()
}

func getTodoData() string {
	text := getUserData("Todo Text: ")
	return text
}


func getUserData(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
