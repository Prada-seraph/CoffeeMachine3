package CoffeeMachine

import (
	"fmt"
	"time"
)


const machineMenu = "Выберите пункт:" +
	"\n1.Коффе" +
	"\n2.Статистика ресурсов" +
	"\n3.Статистика кофе"
const coffeeMenu = "Выберите пункт:" +
	"\n1.Эспрессо" +
	"\n2.Каппучино" +
	"\n3.Мокачино" +
	"\n4.Ристретто" +
	"\n5.Американо" +
	"\n6.Вернуться"

var espStat int
var capStat int
var bamzStat int
var risStat int
var amerStat int


func Mach(m *Machine)  {
	var menuItem int
fmt.Println(machineMenu)
fmt.Scanln(&menuItem)

for true{
	switch menuItem {
	case 1:
		CoffeeItem(m)
		break
	case 2:
		ResStat(m)
		break
	case 3:
		CoffeeStat(m)
		break
	default:
		break
	}
	fmt.Println(machineMenu)
	fmt.Scanln(&menuItem)
}

}
func CoffeeItem(m *Machine){
	var coffeeItem int

	fmt.Println(coffeeMenu)
	fmt.Scanln(&coffeeItem)

	for{
		if m.Coffee<50 || m.Water<50||m.Milk<50 ||m.Cup<1{
			switch coffeeItem {
		case 1:
				fmt.Println("Пополнение ресурсов...\n")
				time.Sleep(5* time.Second)
				m.ReplenishRes()
				espStat++
				fmt.Println("Приготовление...\n")
				time.Sleep(3* time.Second)
				m.CoffeeGive(30, 50, 0, 50, 1)
				fmt.Print("Коффе готово\n")
			break

		case 2:
				fmt.Println("Пополнение ресурсов...\n")
				time.Sleep(5 * time.Second)
				m.ReplenishRes()
				capStat++
				fmt.Println("Приготовление...\n")
				time.Sleep(3 * time.Second)
				m.CoffeeGive(30, 50, 50, 100, 1)
				fmt.Print("Коффе готово\n")
			break

		case 3:
				fmt.Println("Пополнение ресурсов...\n")
				time.Sleep(5 * time.Second)
				m.ReplenishRes()
				bamzStat++
				fmt.Println("Приготовление...\n")
				time.Sleep(3* time.Second)
				m.CoffeeGive(50, 50, 50, 75, 1)
				fmt.Print("Коффе готово\n")
			break

		case 4:
				fmt.Println("Пополнение ресурсов...\n")
				time.Sleep(5 * time.Second)
				m.ReplenishRes()
				risStat++
				fmt.Println("Приготовление...\n")
				time.Sleep(3 * time.Second)
				m.CoffeeGive(30, 30, 0, 50, 1)
				fmt.Print("Коффе готово\n")
			break

		case 5:
				fmt.Println("Пополнение ресурсов...\n")
				time.Sleep(5*time.Second)
				m.ReplenishRes()
				amerStat++
				fmt.Println("Приготовление...\n")
				time.Sleep(3*time.Second)
				m.CoffeeGive(30, 100, 0, 60, 1)
				fmt.Print("Коффе готово\n")
			break

		case 6:
			Mach(m)
			break
		default:
			fmt.Println("Такого пункта нет в списке")
		}
		}
		fmt.Println(coffeeMenu)
		fmt.Scanln(&coffeeItem)
	}
}

func(m *Machine) CoffeeGive(coffee, water,milk,money,cup int)  {
	m.Coffee-=coffee
	m.Water-=water
	m.Milk-=milk
	m.Money+=money
	m.Cup-=cup
	//fmt.Printf("Coffee: %d, Watter: %d, Milk: %d, Money: %d, Cup: %d",m.Coffee, m.Water,m.Milk,m.Money,m.Cup)
}

func(m *Machine) ReplenishRes(){
	m.Coffee+=5000
	m.Water+=5000
	m.Milk+=5000
	m.Cup+=500
}

func ResStat(m *Machine){
	fmt.Printf("\nКоффе: %d,\nВода: %d,\nМолоко: %d,\nДеньги: %d,\nСтаканы: %d",m.Coffee,m.Water,m.Milk,m.Money,m.Cup)
	fmt.Println()
	Mach(m)
}

func CoffeeStat(m *Machine){
	fmt.Printf("\nЭспрессо:%d,\nКаппучино:%d," +
		"\nМокачино:%d,\nРистретто:%d,\nАмерикано:%d",espStat,capStat,bamzStat,risStat,amerStat)
	fmt.Println()
	Mach(m)
}