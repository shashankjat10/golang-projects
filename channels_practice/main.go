package main

import (
	"fmt"
	"time"
)

func main() {
	inputChan := make(chan string)

	// read use input
	readUserInput(inputChan)

	// listen to user input
	for input := range inputChan {
		fmt.Printf("The input was : %s", input)
	}

	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	ch <- "FNH FOREVER!"
	// 	ch <- "Greatt"
	// }()
	time.Sleep(2 * time.Second)
}
