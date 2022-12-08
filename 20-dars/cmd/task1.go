package main

import (
	"fmt"
)

type Sweets interface {
	GetSugar() int
	GetWeight() int
}

type Marmilad struct {
	sugar  int
	weight int
}

func (m Marmilad) GetSugar() int {
	return m.sugar
}

func (m Marmilad) GetWeight() int {
	return m.weight
}

func main() {
	var shirinlik Sweets = Marmilad{sugar: 5, weight: 10}
	fmt.Println(shirinlik.GetSugar())
	fmt.Println(shirinlik.GetWeight())
}
