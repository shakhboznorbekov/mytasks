package main

import "fmt"

// 1-usul// higher order function
// func main() {

// 	//Anonym function

// 	var square = func(a int) int {
// 		return a * a
// 	}

// 	fmt.Println(square(4))
// 	fmt.Println(square(9))
// }

// 2-func
func add(a float64, b float64) float64 {
	return a + b
}

func sub(a float64, b float64) float64 {
	return a - b
}

func build(x func(float64, float64) float64) float64 {
	return x(4, 5)
}

func main() {
	fmt.Println(build(add))
	fmt.Println(build(sub))
}
