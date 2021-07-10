package CoffeeMachine

import "CoffeeMachine3/CoffeeModel"

type Machine struct {
	Money  int
	Coffee int
	Milk   int
	Water  int
	Cup    int

	CoffeeTypes map[string]CoffeeModel.Coffee
}

func NewMachine(money, coffee, milk,water,cup int) *Machine{
	return &Machine{
		Money: money,
		Coffee: coffee,
		Milk: milk,
		Water: water,
		Cup: cup,
	}
}

func (mach *Machine) СoffeeType(){
	var CoffeeTypes = make(map[string]CoffeeModel.Coffee)
	CoffeeTypes["Эспрессо"] = CoffeeModel.Coffee{}
	CoffeeTypes["Каппучино"] =CoffeeModel.Coffee{}
	CoffeeTypes["Мокачино"]= CoffeeModel.Coffee{}
	CoffeeTypes["Ристретто"] = CoffeeModel.Coffee{}
	CoffeeTypes["Американо"] = CoffeeModel.Coffee{}
	mach.CoffeeTypes= CoffeeTypes
}
