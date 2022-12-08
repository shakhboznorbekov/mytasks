package main

import "fmt"

type Transport struct {
	Name     string
	Engine   string
	Distance int
}

type Car struct {
	Info     Transport
	Caracter string
}

func (t Transport) GetName() string {
	return t.Name
}

func (t Transport) GetEngine() string {
	return t.Engine
}

func (t Transport) GetDistance() int {
	return t.Distance
}

func (c Car) GetCaracter() string {
	return c.Caracter
}

type xususiyat interface {
	GetName() string
	GetEngine() string
	GetDistance() int
	GetCaracter() string
}

func Run(x xususiyat) {
	fmt.Println("Avtomobil: "+x.GetName()+"Quvvati: "+x.GetEngine()+"Bora olish masofasi: ", x.GetDistance(), "Yetib borish xususiyati:"+x.GetCaracter())
}

func main() {
	var car = Car{
		Info: Transport{
			Name:     "audi",
			Engine:   "v8",
			Distance: 120,
		},
		Caracter: "4 ta baloni bor",
	}
	Run(car)
	// Run(Car{Transport{
	// 	Name:     "Mercedes",
	// 	Engine:   "250 ot kuchi",
	// 	Distance: 400,

	// 	Caracter: "Yerda yuradi",
	// }})

}
