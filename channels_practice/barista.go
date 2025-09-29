package main

import (
	"context"
	"fmt"
	"time"
)

type Barista struct {
	Name      string             // name
	OrderTime int                // seconds to take customer order
	Available bool               // whether the barista is available
	Greeting  string             // message to greet the user
	Cancel    context.CancelFunc // cancel function to stop the barista from working
}

func NewBarista(name string, orderTime int, greeting string) *Barista {
	return &Barista{
		Name:      name,
		OrderTime: orderTime,
		Greeting:  greeting,
	}
}

// Add the barista to barista pool
func (b *Barista) AddToPool(context context.Context, customerChan chan *Customer, orderChan chan *Order) {
	go func() {
		for {
			select {
			case <-context.Done(): // cancel signal
				fmt.Printf("%s has stopped working.\n", b.Name)
				return
			case customer, ok := <-customerChan:
				if !ok {
					// channel closed, exit all workers
					return
				}
				fmt.Printf("%s, %s!\nI'll take %d seconds to process your order.\nPlease wait.\n",
					b.Greeting, customer.Name, b.OrderTime)
				time.Sleep(time.Duration(b.OrderTime) * time.Second)

				orderNumberCounter++
				fmt.Printf("\n%s, your order token is : %d. Please wait while we prepare your %s.\n",
					customer.Name, orderNumberCounter, customer.Order)

				orderChan <- NewOrder(orderNumberCounter, customer.Order)
				customer.OrderId = orderNumberCounter
			}
		}
	}()
	fmt.Printf("%s is now working at the counter.\n", b.Name)
}
