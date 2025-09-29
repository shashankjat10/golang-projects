package main

type Order struct {
	Id     int
	Amount float32
	Item   string
}

func NewOrder(id int, item string) *Order {
	return &Order{
		Id:   id,
		Item: item,
	}
}
