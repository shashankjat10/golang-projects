package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// assignCommands function takes in respective channels
// and then passes the data correctly
func handleCommands(inputChan chan string, customerChan chan *Customer, orderChan chan *Order) {
	// listen to user input
	for input := range inputChan {
		// fmt.Printf("The input was : %s\n", input)
		data := strings.Split(input, ",")
		command := data[0]
		switch command {
		case CommandAddCustomer:
			handleCustomerAdditionCommand(data, customerChan)
		case CommandAddBarista:
			handleBaristaAdditionCommand(data, customerChan, orderChan)
		case CommandRemoveBarista:
			handleBaristaRemovalCommand(data)
		case CommandAddMachine:
			handleMachineAdditionCommand(data, orderChan)
		case CommandRemoveMachine:
			handleMachineRemovalCommand(data)
		case CommandExit:
			close(inputChan)
			close(customerChan)
			close(orderChan)
			fmt.Println("Shutting down coffee shop")
		default:
			fmt.Println("Invalid command")
		}
	}
}

// handleCustomerAdditionCommand handles the custom addition
func handleCustomerAdditionCommand(command []string, customerChan chan *Customer) error {
	newCustomer, err := getCustomerFromCommand(command)
	if err != nil {
		fmt.Printf("Error while handling customer command : %v\n%s\n", err, CommandHelpAddCustomer)
		return err
	}
	customerChan <- newCustomer
	return nil
}

// handleBaristaAdditionCommand handles a barista addition
func handleBaristaAdditionCommand(command []string, customerChan chan *Customer, orderChan chan *Order) error {
	newBarista, err := getBaristaFromCommand(command)
	if err != nil {
		fmt.Printf("Error while handling barista+ command : %v\n%s\n", err, CommandHelpAddBarista)
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	newBarista.Cancel = cancel
	availableBaristas = append(availableBaristas, newBarista)
	newBarista.AddToPool(ctx, customerChan, orderChan)
	return nil
}

// handleBaristaRemovalCommand handles a barista removal
func handleBaristaRemovalCommand(command []string) error {
	baristaIndex, err := getAvailableBaristaIndexFromCommand(command)
	if err != nil {
		fmt.Printf("Error while handling barista- command : %v\n%s\n", err, CommandHelpRemoveBarista)
		return err
	}
	availableBaristas = append(availableBaristas[:baristaIndex], availableBaristas[baristaIndex+1:]...)
	return nil
}

// handleMachineAdditionCommand handles a machine addition
func handleMachineAdditionCommand(command []string, orderChan chan *Order) error {
	newMachine, err := getMahineFromCommand(command)
	if err != nil {
		fmt.Printf("Error while handling machine+ command : %v\n%s\n", err, CommandHelpAddMachine)
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	newMachine.Cancel = cancel
	availableMachines = append(availableMachines, newMachine)
	newMachine.AddToPool(ctx, orderChan)
	return nil
}

// handleMachineRemovalCommand handles a machine removal
func handleMachineRemovalCommand(command []string) error {
	machineIndex, err := getAvailableMachineIndexFromCommand(command)
	if err != nil {
		fmt.Printf("Error while handling machine- command : %v\n%s\n", err, CommandHelpRemoveMachine)
		return err
	}
	availableMachines = append(availableMachines[:machineIndex], availableMachines[machineIndex+1:]...)
	return nil
}

// getCustomerFromCommand gets a new customer from a customer command split
func getCustomerFromCommand(command []string) (*Customer, error) {
	if len(command) != 4 {
		return nil, errors.New("invalid customer command")
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
	if len(command) != 4 {
		return nil, errors.New("invalid barista+ command")
	}
	name := command[1]
	orderTime, err := strconv.Atoi(command[2])
	if err != nil {
		return nil, fmt.Errorf("invalid order time, must be integer")
	}
	greeting := command[3]

	return NewBarista(name, orderTime, greeting), nil
}

// getAvailableBaristaFromCommand get the barista name from command, and then
// gets that particular barista from the working baristas
func getAvailableBaristaIndexFromCommand(command []string) (int, error) {
	if len(command) != 2 {
		return -1, errors.New("invalid barista- command")
	}
	name := command[1]
	if availableBaristas == nil {
		return -1, errors.New("no barista working at the moment")
	}
	for i, barista := range availableBaristas {
		if barista.Name == name {
			return i, nil
		}
	}
	return -1, errors.New("could not find barista to remove")
}

// getMahineFromCommand gets a new coffee machine from a command
func getMahineFromCommand(command []string) (*CoffeeMachine, error) {
	if len(command) != 3 {
		return nil, errors.New("invalid machine+ command")
	}
	name := command[1]
	prepTime, err := strconv.Atoi(command[2])
	if err != nil {
		return nil, fmt.Errorf("invalid prep time, must be integer")
	}

	return NewCoffeeMachine(prepTime, name), nil
}

// getAvailableMachinesFromCommand gets the machine name from command, and then
// gets that particular machine from the working machines
func getAvailableMachineIndexFromCommand(command []string) (int, error) {
	if len(command) != 2 {
		return -1, errors.New("invalid machine- command")
	}
	name := command[1]
	if availableMachines == nil {
		return -1, errors.New("no working machines at the moment")
	}
	for i, machine := range availableMachines {
		if machine.Name == name {
			return i, nil
		}
	}
	return -1, errors.New("could not find machine to remove")
}
