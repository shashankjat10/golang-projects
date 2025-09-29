package main

import (
	"errors"
	"strconv"
)

// assignCommands function takes in respective channels
// and then passes the data correctly
func handleCommands(inputChan chan string) {

}

// get a new customer from a customer command split
func getCustomerFromCommand(command []string) (*Customer, error) {
	if len(command) != 4 {
		return nil, errors.New("Please enter a valid customer command")
	}

	// take apart the customer entry data
	name := command[1]
	order := command[2]
	consumptionTime, err := strconv.Atoi(command[3])
	if err != nil {
		return nil, err
	}

	return NewCustomer(name, order, consumptionTime), nil
}
