package main

import (
	"fmt"
)

const (
	ADD = iota
	SUB
	DIV
	MULT
)

func calc(a, b, c int) int {
	if c == ADD {
		return a + b
	} else if c == SUB {
		return a - b
	} else if c == DIV {
		return a / b
	} else if c == MULT {
		return a * b
	}
	return 0
}

func main() {
	// const SUB= 1
	// var x int =0
	// x=SUB
	// SUB=x

	fmt.Println(calc(10, 5, ADD))
}
