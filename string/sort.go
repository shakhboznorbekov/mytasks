package main

import (
	"fmt"
)

// func main() {
// 	var nums = []int{3, 5, 1, 2, 70, 23, 8}
// 	// var count int

// 	// for i := 0; i < len(nums); i++ {
// 	// 	for j := 1; j < len(nums)-i; j++ {
// 	// 		if nums[j-1] > nums[j] {
// 	// 			nums[j], nums[j-1] = nums[j-1], nums[j]
// 	// 		}
// 	// 		count++
// 	// 	}
// 	// }
// 	// fmt.Println(nums)
// 	// fmt.Println(count)

// 	sort.Ints(nums)
// 	sort.Slice(nums, func(a, b int) bool {
// 		return nums[a] < nums[b]
// 	})
// 	fmt.Println(nums)
// }

type User struct {
	name string
	age  int
}

func main() {
	var user User = User{
		name: "Shaxboz",
		age:  23,
	}

	fmt.Println(user)
	// sort(user, AGE)
	// sort(user, NAME)
}
