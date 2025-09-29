package main

type Order struct {
	Id     int
	Amount float32
}

func NewOrder(id int) *Order {
	return &Order{
		Id: id,
	}
}
