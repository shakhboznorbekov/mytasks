package main

import (
	"fmt"
	"strings"
)

func swap(x, y *int) {
	*x = *x * 2
	*y = *y * 2
}

func kvadrat(nums []int) {
	for index, _ := range nums {
		nums[index] = nums[index] * nums[index]
	}
}

func toUpper(letters map[int]string) {
	for key, value := range letters {
		letters[key] = strings.ToUpper(value)
	}
	delete(letters, 3)
}

func main() {
	// call by value
	// n:=10
	// m:=n

	// reference
	// n :=10
	// m:= &n

	// *m=20

	// fmt.Println("m:%T%v, n:%T %v\n", m, m,n,n)

	// fmt.Println(*m, n)

	// a, b := 10,15

	// swap(&a, &b)

	// fmt.Println(a, b)

	nums := []int{3, 4, 5, 6}
	letters := map[int]string{
		1: "a",
		2: "b",
		3: "c",
		4: "d",
	}

	kvadrat(nums)

	toUpper(letters)

	fmt.Println(nums)
	fmt.Println(letters)
}
