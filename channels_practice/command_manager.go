package main

import (
	"errors"
	"strconv"
)

// assignCommands function takes in respective channels
// and then passes the data correctly
func handleCommands(inputChan chan string) {

}

// getCustomerFromCommand gets a new customer from a customer command split
func getCustomerFromCommand(command []string) (*Customer, error) {
	if len(command) != 4 {
		return nil, errors.New("Invalid custom command")
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

// getBaristaFromCommand gets a new barista from a command
func getBaristaFromCommand(command []string) (*Barista, error) {
	if len(command) != 3 {
		return nil, errors.New("Invalid barista+ command")
	}
	name := command[1]
	orderTime, err := strconv.Atoi(command[2])
	if err != nil {
		return nil, err
	}
	greeting := command[3]

	return NewBarista(name, orderTime, greeting), nil
}

// getAvailableBaristaFromCommand get the barista name from command, and then
// gets that particular barista from the working baristas
func getAvailableBaristaFromCommand(command []string) (*Barista, error) {
	if len(command) != 2 {
		return nil, errors.New("Invalid barista- command")
	}
	name := command[1]
	if availableBaristas == nil {
		return nil, errors.New("No barista working at the moment")
	}
	for _, barista := range availableBaristas {
		if barista.Name == name {
			return barista, nil
		}
	}
	return nil, errors.New("Could not find barista to remove.")
}
