package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// read user input through the terminal
func readUserInput(inputChan chan string) {
	input := ""
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter command : ")
		if scanner.Scan() { // reads until newline
			input = scanner.Text()
		}
		inputChan <- input
		time.Sleep(500 * time.Millisecond)
	}
}
