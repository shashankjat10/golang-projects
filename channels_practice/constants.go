package main

const (
	CommandAddCustomer   = "customer"
	CommandAddBarista    = "barista+"
	CommandRemoveBarista = "barista-"
	CommandAddMachine    = "machine+"
	CommandRemoveMachine = "machine-"
	CommandExit          = "exit"
)

const (
	CommandHelpAddCustomer   = "Adding a new Customer : \ncustomer,<CustomerName(String)>,<OrderName(String)>,<ConsumptionTime(Integer)>"
	CommandHelpAddBarista    = "Adding a new barista: \nbarista+,<Name(String)>,<OrderHandlingTime(Integer)>,<Greeting(String)>"
	CommandHelpRemoveBarista = "Removing a barista: \nbarista-,<Name(String)>"
	CommandHelpAddMachine    = "Adding a new machine: \nmachine+,<Name(String)>,<PrepTime(Integer)>"
	CommandHelpRemoveMachine = "Removing a machine: \nmachine-,<Name(String)>"
	CommandHelpExit          = "Shutting down the coffee shop: \nexit"
)
