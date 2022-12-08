package main

import (
	"fmt"
)

func count(name string, l rune) int {
	var c int
	for _, v := range name {
		if v == l {
			c++
		}
	}
	return c
}

func main() {
	var name = "Qo'yhokiz"
	var letter1 rune = 'Q'
	var letter2 rune = 'y'
	var letter3 rune = 'h'
	var letter4 rune = 'k'
	var letter5 rune = 'i'
	var letter6 rune = 'z'
	var letter7 rune = '\''

	c := count(name, letter1)
	a := count(name, letter2)
	b := count(name, letter3)
	d := count(name, letter4)
	s := count(name, letter5)
	e := count(name, letter6)
	h := count(name, letter7)
	fmt.Println(name)

	fmt.Printf(string(letter1), c)
	fmt.Println(string(letter2), a)
	fmt.Println(string(letter3), b)
	fmt.Println(string(letter4), d)
	fmt.Println(string(letter5), s)
	fmt.Println(string(letter6), e)
	fmt.Println(string(letter7), h)
}
