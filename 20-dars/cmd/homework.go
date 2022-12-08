package main

import (
	"fmt"
)

type Car interface {
	GetEngine() string
	GetInfo() string
	GetSpeed() int
}

type Audi struct {
	Engine string
	Info   string
	Speed  int
}

func (a Audi) GetEngine() string {
	return a.Engine
}

func (a Audi) GetInfo() string {
	return a.Info
}

func (a Audi) GetSpeed() int {
	return a.Speed
}

type Cadillac struct {
	Engine string
	Info   string
	Speed  int
}

func (c Cadillac) GetEngine() string {
	return c.Engine
}

func (c Cadillac) GetInfo() string {
	return c.Info
}

func (c Cadillac) GetSpeed() int {
	return c.Speed
}
func main() {
	var cars Car = Audi{Engine: "220 ot kuchi", Info: "Audi X6 2022yilda ishlab chiqarilgan", Speed: 250}
	fmt.Println(cars.GetEngine())
	fmt.Println(cars.GetInfo())
	fmt.Println(cars.GetSpeed())

	var avto Car = Cadillac{Engine: "300 ot kuchi", Info: "Cadillac Escalado 2022yilda ishlab chiqarilgan", Speed: 270}
	fmt.Println(avto.GetEngine())
	fmt.Println(avto.GetInfo())
	fmt.Println(avto.GetSpeed())

}
