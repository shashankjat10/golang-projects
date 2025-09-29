package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// read user input through the terminal
func readUserInput(inputChan chan string) {
	fmt.Printf("Use these commands to run the coffee shop: \n%s\n%s\n%s\n%s\n%s\n%s\n",
		CommandHelpAddCustomer, CommandHelpAddBarista, CommandHelpRemoveBarista, CommandHelpAddMachine, CommandHelpRemoveMachine, CommandHelpExit)
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
