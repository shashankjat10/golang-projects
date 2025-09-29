package main

import (
	"context"
	"fmt"
	"time"
)

type CoffeeMachine struct {
	PrepTime      int                // seconds
	Name          string             // title
	Charge        float32            // current battery percentage
	DischargeRate float32            // discharge rate
	Cancel        context.CancelFunc // cancel function to stop the machine from working
}

func NewCoffeeMachine(prepTime int, name string) *CoffeeMachine {
	return &CoffeeMachine{
		PrepTime: prepTime,
		Name:     name,
	}
}

// Add the coffee machine to machine pool
func (cm *CoffeeMachine) AddToPool(context context.Context, orderChan chan *Order) {
	go func() {
		select {
		case <-context.Done():
			fmt.Printf("%s has stopped working.\n", cm.Name)
			return
		case order, ok := <-orderChan:
			if !ok {
				// channel closed
				return
			}
			fmt.Printf("\n%s is prepping the order %d, %s. Will take %d seconds.\n",
				cm.Name, order.Id, order.Item, cm.PrepTime)
			time.Sleep(time.Duration(cm.PrepTime) * time.Second)
			fmt.Printf("\n%d, %s is prepared and ready.\n", order.Id, order.Item)
		}
	}()
	fmt.Printf("%s is now running.\n", cm.Name)
}
