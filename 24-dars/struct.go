package main

import (
	"fmt"
)

type Family struct {
	Adult
	Child
	Parent
}

type Parent struct {
	Name       string
	Age        int
	Location   string
	Profession string
}

type Child struct {
	Age      int
	Name     string
	Location string
	School   string
}

type Adult struct {
	Single  bool
	Pensiya float64
}

func main() {
	var family Family = Family{
		Adult: Adult{
			Single:  false,
			Pensiya: 210000,
		},
		Child: Child{
			Age:      21,
			Name:     "Farxod",
			Location: "Toshkent",
			School:   "32-maktab",
		},
		Parent: Parent{
			Name:       "Davron",
			Age:        50,
			Location:   "Toshkent",
			Profession: "Enginer",
		},
	}

	fmt.Println(family)

}
