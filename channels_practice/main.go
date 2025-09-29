package main

var availableBaristas []*Barista
var availableMachines []*CoffeeMachine
var orderNumberCounter int = 0

func main() {
	inputChan := make(chan string, 1)

	// read use input
	go readUserInput(inputChan)

	// customer queue channel
	customerQueueChan := make(chan *Customer, 10)

	// order channel
	orderChan := make(chan *Order, 10)

	// handle the input commands
	handleCommands(inputChan, customerQueueChan, orderChan)
}
