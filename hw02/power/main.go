package main

import "fmt"

type Engine struct {
	Power string
}

type Radio struct {
	Power string
}

type Car struct {
	Engine
	Radio
}

func main() {
	car := Car{}

	// без явного указания
	// car.Power = "on"
	// hw02/power/main.go:19:6: ambiguous selector car.Power

	// явное указание
	car.Engine.Power = "on"
	car.Radio.Power = "off"

	fmt.Println(car)
}
