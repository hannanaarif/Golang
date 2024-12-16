package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello! ", phrase)
	doneChan <- true
}

func SlowGreet(phrase string, doneChan chan bool) {
	time.Sleep(5 * time.Second)
	fmt.Println("Hello pharse from slow greet!", phrase)
	doneChan <- true
}

func main() {
	done := make(chan bool)

	go greet("Nice to meet you!", done)
	go greet("How are you?", done)
	go SlowGreet("How...are...you", done)
	go greet("I hope you are having amazing days", done)
	<-done
	<-done
	<-done
	<-done
}
