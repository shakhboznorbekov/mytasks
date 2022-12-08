package main

import (
	"fmt"
)

// 1-masala
// func main() {
// 	for i := 35; i < 87; i++ {
// 		if i%7 == 1 || i%7 == 2 || i%7 == 5 {
// 			fmt.Println(i)
// 		}
// 	}
// }

// 2-masala
// func main() {
// 	var son, son2 int = 111, 0
// 	for i := son; i != 0; i /= 10 {
// 		son2 = (son2 * 10) + i%10
// 	}
// 	if son == son2 {
// 		fmt.Println("Polindrome son")
// 	} else {
// 		fmt.Println("Polindrome son emas")
// 	}
// }

// 3-masala
// func main() {
// 	var result []int = []int{}
// 	result = rec(12345)
// 	fmt.Println(result)

// }

// func rec(num int) []int {
// 	if num == 0 {
// 		return []int{}
// 	}
// 	if (num%10)%2 == 1 {
// 		res := rec(num / 10)
// 		res = append(res, num%10)
// 		return res
// 	}
// 	return rec(num / 10)
// }

// // 4-masala
// func main() {
// 	var son, result int = 818, 0
// 	for i := 1; i < son; i++ {
// 		if son%i == 0 {
// 			result += i
// 		}
// 	}
// 	if result == son {
// 		fmt.Println("Mukammal")
// 	} else {
// 		fmt.Println("Mukammal emas")
// 	}
// }

// 5-masala
// func main() {
// 	var son, result int = 915457, 0
// 	for son > 9 {
// 		result = son % 10
// 		son /= 10
// 	}
// 	fmt.Println(result)
// }

// 3-masala
// func main() {
// 	// var son int = 1234567
// 	// for son > 0 {
// 	// 	if son%2 != 0 {
// 	// 		result = append(result, son%10)
// 	// 	}
// 	// 	son /= 10
// 	// }
// 	// var max = result[0]
// 	// for _, value := range result {
// 	// 	for _, val := range result{
// 	// 		if val>
// 	// 	}
// 	// }

// 	fmt.Println(result)

// }
