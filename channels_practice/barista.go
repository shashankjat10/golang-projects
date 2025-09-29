package main

import (
	"fmt"
	"time"
)

type Barista struct {
	Name      string // name
	OrderTime int    // seconds to take customer order
	Available bool   // whether the barista is available
	Greeting  string // message to greet the user
}

func NewBarista(name string, orderTime int, greeting string) *Barista {
	return &Barista{
		Name:      name,
		OrderTime: orderTime,
		Greeting:  greeting,
	}
}

// Add the barista to barista pool
func (b *Barista) AddToPool(customerChan chan *Customer, orderChan chan *Order) {
	go func() {
		for customer := range customerChan {
			fmt.Printf("%s, %s!\nI'll take %d seconds to process your order.\nPlease wait.\n",
				b.Greeting, customer.Name, b.OrderTime)
			time.Sleep(time.Duration(b.OrderTime) * time.Second)
			// increase the order counter
			orderNumberCounter++
			fmt.Printf("\n%s, your order token is : %d. Please wait while we prepare your %s.\n",
				customer.Name, orderNumberCounter, customer.Order)
			// pass the order to order channel
			orderChan <- NewOrder(orderNumberCounter, customer.Order)
			// update the customer object with the order id
			customer.OrderId = orderNumberCounter
		}
	}()
	fmt.Printf("%s is now working at the counter.\n", b.Name)
}
