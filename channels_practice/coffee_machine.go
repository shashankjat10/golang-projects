package main

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
