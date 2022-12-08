package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var nums []int

	for i := 0; i < 20; i++ {
		nums = append(nums, rand.Intn(20))
	}
	lens := len(nums)
	for i := 0; i < lens; i++ {
		for j := i + 1; j < lens; j++ {
			if nums[i] == nums[j] {
				fmt.Println(j)
				return
			}
		}
	}
	fmt.Println(nums)

}
