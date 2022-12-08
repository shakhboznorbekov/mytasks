package main

import (
	"fmt"
)

//1-masala
// func direction(a interface{}) {
// 	if a == 0 {
// 		fmt.Println(reflect.TypeOf(a).Name() == "string", "TOP")
// 	} else if a == "TOP" {
// 		fmt.Println(reflect.TypeOf(a).Name() == "int", "0")
// 	} else if a == 1 {
// 		fmt.Println(reflect.TypeOf(a).Name() == "string", "BOTTOM")
// 	} else if a == "BOTTOM" {
// 		fmt.Println(reflect.TypeOf(a).Name() == "int", "1")
// 	} else if a == 2 {
// 		fmt.Println(reflect.TypeOf(a).Name() == "string", "RIGHT")
// 	} else if a == "RIGHT" {
// 		fmt.Println(reflect.TypeOf(a).Name() == "int", "2")
// 	} else if a == 3 {
// 		fmt.Println(reflect.TypeOf(a).Name() == "string", "LEFT")
// 	} else if a == "Left" {
// 		fmt.Println(reflect.TypeOf(a).Name() == "int", "3")
// 	} else {
// 		fmt.Println("true", "")
// 	}
// }

// func main() {

// 	// 	0="Top"
// 	// 	1="BOTTOM"
// 	// 	2="RIGHT"
// 	// 	3="LEFT"

// 	direction(1)
// 	direction("Left")
// 	direction(5)

// }

// 2-masala
// func culc(a int, b int) (func() int, func() int) {
// 	return func() int { return a + b }, func() int { return a - b }
// }

// func main() {
// 	plus, minus := culc(30, 10)
// 	fmt.Println(plus())
// 	fmt.Println(minus())
// }

// 3-masala
func MaxAndMin(nums ...int) (int, int) {
	var max = nums[0]
	var min = nums[0]
	for _, val := range nums {
		if max < val {
			max = val
		}
		if min > val {
			min = val
		}
	}
	return max, min
}

func main() {
	max, min := MaxAndMin(1, 2, 3, 4, 5, 10, -1, -3)
	fmt.Println(max)
	fmt.Println(min)
}
