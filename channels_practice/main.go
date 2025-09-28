package main

import (
	"fmt"
)

var coffeeMachines int = 3
var baristas int = 2

func main() {
	inputChan := make(chan string)

	// read use input
	go readUserInput(inputChan)

	// customer queue channel
	// customerQueueChan := make(chan string, 10)

	// // barista channel
	// baristaChan := make(chan string, 2)

	// listen to user input
	for input := range inputChan {
		fmt.Printf("The input was : %s\n", input)

		switch input {

		}
	}

	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	ch <- "FNH FOREVER!"
	// 	ch <- "Greatt"
	// }()
	// time.Sleep(2 * time.Second)
}
