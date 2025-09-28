package main

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
