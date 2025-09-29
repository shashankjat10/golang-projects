package main

var coffeeMachines int = 3
var baristas int = 2
var availableBaristas []*Barista
var availableMachines []*CoffeeMachine

func main() {
	inputChan := make(chan string)

	// read use input
	go readUserInput(inputChan)

	// customer queue channel
	customerQueueChan := make(chan *Customer, 10)

	// handle the input commands
	handleCommands(inputChan, customerQueueChan)

	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	ch <- "FNH FOREVER!"
	// 	ch <- "Greatt"
	// }()
	// time.Sleep(2 * time.Second)
}
