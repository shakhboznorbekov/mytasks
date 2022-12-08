package main

import "fmt"

func sum(nums ...int) int {
	var result int
	for _, val := range nums {
		result += val
	}
	return result
}

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(sum(nums...))
}
