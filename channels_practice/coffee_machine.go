package main

import (
	"fmt"
	"time"
)

type CoffeeMachine struct {
	PrepTime      int     // seconds
	Name          string  // title
	Charge        float32 // current battery percentage
	DischargeRate float32 // discharge rate
}

func NewCoffeeMachine(prepTime int, name string) *CoffeeMachine {
	return &CoffeeMachine{
		PrepTime: prepTime,
		Name:     name,
	}
}

// Add the coffee machine to machine pool
func (cm *CoffeeMachine) AddToPool(orderChan chan *Order) {
	go func() {
		for order := range orderChan {
			fmt.Printf("\n%s is prepping the order %d, %s. Will take %d seconds.\n",
				cm.Name, order.Id, order.Item, cm.PrepTime)
			time.Sleep(time.Duration(cm.PrepTime) * time.Second)
			fmt.Printf("\n%d, %s is prepared and ready.\n", order.Id, order.Item)
		}
	}()
	fmt.Printf("%s is now running.\n", cm.Name)
}
