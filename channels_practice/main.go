package main

import (
	"fmt"
	"strings"
)

var coffeeMachines int = 3
var baristas int = 2

const (
	CommandAddCustomer   = "customer"
	CommandAddBarista    = "barista+"
	CommandRemoveBarista = "barista-"
	CommandAddMachine    = "machine+"
	CommandRemoveMachine = "machine-"
	CommandExit          = "exit"
)

func main() {
	inputChan := make(chan string)

	// read use input
	go readUserInput(inputChan)

	// customer queue channel
	// customerQueueChan := make(chan string, 10)

	// listen to user input
	for input := range inputChan {
		fmt.Printf("The input was : %s\n", input)
		data := strings.Split(input, ",")
		command := data[0]
		switch command {
		case CommandAddCustomer:
			fmt.Println("Added a new custom")
		case CommandAddBarista:
			fmt.Println("added a barista")
		case CommandRemoveBarista:
			fmt.Println("Removed a barista")
		case CommandAddMachine:
			fmt.Println("added a machine")
		case CommandRemoveMachine:
			fmt.Println("removed a machine")
		case CommandExit:
			close(inputChan)
			fmt.Println("Shutting down coffee shop")
		default:
			fmt.Println("Invalid command")

		}
	}

	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	ch <- "FNH FOREVER!"
	// 	ch <- "Greatt"
	// }()
	// time.Sleep(2 * time.Second)
}
