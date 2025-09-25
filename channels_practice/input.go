package main

import (
	"bufio"
	"fmt"
	"os"
)

// read user input through the terminal
func readUserInput(inputChan chan string) {
	input := ""
	scanner := bufio.NewScanner(os.Stdin)

	for input != "exit" {
		fmt.Print("Enter command : ")
		if scanner.Scan() { // reads until newline
			input = scanner.Text()
		}
		inputChan <- input
	}
	close(inputChan)
}
