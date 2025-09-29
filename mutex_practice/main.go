package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

var sharedCounter int = 0

func main() {
	fmt.Println(`This program is used to increase a counter whose value is set to 0.
It will increase it a 1000 times using parallel go routines. You can choose 
whether to use mutexe or not. If you want to use mutex, type Y, otherwise type N.`)
	input := ""
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Y/N: ")
	if scanner.Scan() {
		input = scanner.Text()
	}
	switch strings.ToLower(input) {
	case "y":
		incrementWithMutex()
	case "n":
		incrementWithoutMutex()
	default:
		fmt.Println("You are stupid as fuck")
	}
	fmt.Printf("Final value of counter : %d\n", sharedCounter)
}

func incrementWithMutex() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			oldValue := sharedCounter
			sharedCounter = oldValue + 1
			mu.Unlock()
		}()
	}
	wg.Wait()
}

// increment the values without a mutex
func incrementWithoutMutex() {
	var wg sync.WaitGroup
	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}
	wg.Wait()
}

// function to increment
func increment() {
	oldValue := sharedCounter
	sharedCounter = oldValue + 1
}
