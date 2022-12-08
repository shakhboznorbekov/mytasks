package main

import (
	"fmt"
)

// func check(arr []int, son int) (int, bool) {
// 	for index, item := range arr {
// 		if son == item {
// 			return index, true
// 		}
// 	}
// 	return -1, false
// }

// func main() {
// 	var sonlar = []int{1, 2, 3, 4, 5, 6, 7, 8}
// 	son := 10
// 	if resp, ok := check(sonlar, son); ok {
// 		fmt.Println(resp)
// 	} else {
// 		fmt.Println("Sorry", resp)
// 	}
// }

// 2-masala
// func PowFunc(son int, daraja int) float64 {
// 	var result float64 = 1
// 	if daraja > 0 {
// 		for i := 1; i <= daraja; i++ {
// 			result *= float64(son)
// 		}
// 		return float64(result)
// 	} else {
// 		for i := daraja; i <= -1; i++ {
// 			result *= 1 / float64(son)
// 		}
// 		return float64(result)
// 	}
// 	return result

// }

// func main() {
// 	var son int = 2
// 	var daraja int = 4
// 	fmt.Println(PowFunc(son, daraja))

// }

// 3-masala
func sorting(sonlar []int) []int {

	for i := 0; i < len(sonlar); i++ {
		for j := 0; j < len(sonlar)-i-1; j++ {
			if sonlar[j] > sonlar[j+1] {
				sonlar[j], sonlar[j+1] = sonlar[j+1], sonlar[j]
			}
		}
	}
	return sonlar

}

func main() {
	sonlar := []int{1, 2, 3, 7, 8, 4, 5, 9, 15, 14, 11}

	fmt.Println(sorting(sonlar))
}
