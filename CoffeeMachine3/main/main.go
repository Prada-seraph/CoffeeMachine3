package main

import (
	"CoffeeMachine3/CoffeeMachine"
)

func main(){
	cm := CoffeeMachine.NewMachine(0,5000,5000,5000,500)
	cm.СoffeeType()
	CoffeeMachine.Mach(cm)
}