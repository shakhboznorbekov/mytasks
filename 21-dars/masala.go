package main

import (
	"fmt"
	//"time"
)

// func Resultat(a int) func(b int) func(c int) int {
// 	return func(b int) func(c int) int {
// 		return func(c int) int {
// 			return a * b * c
// 		}
// 	}
// }

// func main() {
// 	res := Resultat(2)(3)(4)
// 	fmt.Println(res)
// }

// //2-masala

// func squareSum(input int, ch chan int) {
// 	start := time.Now()
// 	var output int

// 	for input != 0 {
// 		n := input % 10
// 		output += n * n
// 		input /= 10
// 	}

// 	ch <- output
// 	if time.Now().Sub(start) > time.Duration(2) {
// 		panic("jhdsij")
// 	}
// }

// func cubeSum(input int, ch chan int) {

// 	var output int

// 	for input != 0 {
// 		n := input % 10
// 		output += n * n * n
// 		input /= 10
// 	}

// 	ch <- output
// }

// func main() {

// 	square_channel := make(chan int)
// 	cube_channel := make(chan int)

// 	input := 123

// 	go squareSum(input, square_channel)
// 	go cubeSum(input, cube_channel)

// 	x := <-square_channel
// 	y := <-cube_channel

// 	fmt.Println(x, y)

// }

// start :=time.Now()

//	if time.Now().Sub(start) > time.Duration(2) {
//		panic("jhdsij")
//	}

//3-masala

// func check(a int) string {
// 	if a < 0 {
// 		return "kiritilgan son manfiy"
// 	} else if a > 0 {
// 		return "kiritilgan son musbat"
// 	}
// 	return "kiritilgan son 0ga teng"
// }
// func main() {
// 	var a int = -7
// 	fmt.Println(check(a))
// }

//2-usul
// func main() {

// 	a := 2
// 	if a < 0 {
// 		fmt.Println("kiritilgan son manfiy")
// 	} else if a > 0 {
// 		fmt.Println("kiritilgan son musbat")
// 	}else{
// 		fmt.Println("kiritilgan son 0ga teng")
// 	}
// }

// 3-usul
// func main() {

// 	a := 2
// 	if a < 0 {
// 		fmt.Println("kiritilgan son manfiy")
// 	} else if a > 0 {
// 		fmt.Println("kiritilgan son musbat")
// 	}
// 	switch a {
// 	case 0:
// 		fmt.Println("0")
// 	}
// }

// 4-masala
func dublikat(list []int) []int {
	var b []int
	for _, val := range list {
		for _, val2 := range list {
			if val != val2 {
				b = append(b, val)
				break
			}
		}
	}
	return b
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 7, 7, 8, 8, 9, 1, 2}
	fmt.Println(dublikat(a))
}
