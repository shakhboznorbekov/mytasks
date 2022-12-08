package main

import "fmt"

// func kop(arr []int) []int {
// 	length := len(arr)
// 	arr2 := make([]int, length)

// 	for i := 0; i < length; i++ {
// 		if i-1 >= 0 && i+1 < length {                 //SHARTI
// 			arr2[i] = arr[i-1] * arr[i+1]
// 		}                                          //INPUT -->  {1, 2, 3, 4, 5, 6}

// 		if i == 0 {                                //OUTPUT --> [2 3 8 15 24 30]
// 			arr2[i] = arr[i] * arr[i+1]
// 		}

// 		if i == length-1 {
// 			arr2[i] = arr[i] * arr[i-1]
// 		}
// 	}
// 	return arr2

// }

// func main() {
// 	var arr []int = []int{1, 2, 3, 4, 5, 6}

// 	fmt.Println(kop(arr))
// }

//HOMEWORK

//5-MASALA

// INPUT: apple      INPUT: 12345

// OUTPUT:            OUTPUT:
// apple              12345
// eappl              51234
// leapp              54123
// pleap              34512
// pplea              23451

// func lenLastword(s string) {
// 	lenS := len(s)
// 	// var a string
// 	for i := lenS - 1; i >= 0; i-- {
// 		fmt.Printf("%s", string(s[i]))
// 		fmt.Printf("%v", s[:i])
// 		s = string(s[i]) + s[:i]
// 		fmt.Println()

// 	}
// }
// func main() {

// 	intput := "apple"
// 	lenLastword(intput)

// }

func main() {

	intput := "apple"
	lenS := len(intput)

	for lenS > 0 {
		fmt.Printf("%s", string(intput[lenS-1]))
		fmt.Printf("%v\n", intput[:lenS-1])
		intput = string(intput[lenS-1]) + intput[:lenS-1]
		fmt.Println()
		lenS--

	}
}

//6-masala

// Write a program in Go to rearrange an slice in such an order
// that– smallest, largest, 2nd smallest, 2nd largest and on.

// Expected Output:
// The given array is:
// 5 8 1 4 2 9 3 7 6

// The new array is:
// 1 9 2 8 3 7 4 6 5

// func main() {
// 	var arr []int = []int{5, 8, 1, 4, 2, 9, 3, 7, 6}
// 	length := len(arr)
// 	for i := 0; i < length; i++ {
// 		for j := 0; j < length; j += 2 {
// 			if arr[i] < arr[j] && i%2 == 0 {
// 				arr[i], arr[j] = arr[j], arr[i]
// 			}
// 		}
// 		for k := 1; k < length; k += 2 {
// 			if arr[i] > arr[k] && k%2 != 0 {
// 				arr[i], arr[k] = arr[k], arr[i]
// 			}
// 		}
// 	}
// 	fmt.Println(arr)

// }
