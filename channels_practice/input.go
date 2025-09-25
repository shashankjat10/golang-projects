package main

import "fmt"

// read user input through the terminal
func readUserInput(inputChan chan string) {
	input := ""
	for input != "exit" {
		fmt.Scanln(&input)
		inputChan <- input
	}
	close(inputChan)
}
