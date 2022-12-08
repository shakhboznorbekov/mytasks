package main

import (
	"fmt"
)

type AbstractCar interface {
	GetName() string
}

type Audi struct {
	name string
}

func (a Audi) GetName() string {
	return a.name
}

type BMW struct {
	name string
}

func (b BMW) GetName() string {
	return b.name
}

func main() {
	var car AbstractCar = Audi{name: "Audi"}
	fmt.Println(car.GetName())
}
