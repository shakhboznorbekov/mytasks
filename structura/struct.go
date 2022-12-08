package main

import "fmt"



type komputer struct {
	CPU string         //field
	Memory string
	RAM string
	
}

// func (c komputer) getName() string{
// 	return string(c.OS)
// }



func main(){

	var acer komputer 
	acer= komputer{
		CPU: "aceeer",
		Memory: "HDD512GB",
		RAM: "AMDRYZEN"

	}

	// fmt.Println(acer.GetName())
	// fmr.Println(acer.OS.GetName())
	fmt.Println(acer.CPU)
	fmt.Println(acer.Memory)
	fmt.Println(acer.RAM)
}
