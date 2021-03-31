package main

import "fmt"

func greeting(myChannel chan string) {
	myChannel <- "hi"
}

func main() {
	myChanel := make(chan string)
	go greeting(myChannel)
	receivedValue := <-myChanel
	fmt.Println(receivedValue)
}
