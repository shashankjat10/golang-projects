package main

type Customer struct {
	Name            string // customer name
	Order           string // order given
	ConsumptionTime int    // time to consume ordered product in seconds
	OrderId         int    // id of the order
}

func NewCustomer(name string, order string, time int) *Customer {
	return &Customer{
		Name:            name,
		Order:           order,
		ConsumptionTime: time,
	}
}
