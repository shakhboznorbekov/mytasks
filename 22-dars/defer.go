// package main

// import (
// 	"fmt"
// )

// func add(a, b int) int {
// 	return a + b
// }

// func main() {

// 	defer func() {
// 		a := add(3, 4)
// 		fmt.Println(a)
// 	}()

// 	i := 0
// 	defer fmt.Println("A", i)

// 	i = 1

// 	defer fmt.Println("B", i)

// 	i = 2

// 	fmt.Println("C", i)
// }
