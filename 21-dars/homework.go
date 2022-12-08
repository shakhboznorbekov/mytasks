package main

import (
	"fmt"
)

type Product struct {
	name  string
	price float64
}

type Client struct {
	name          string
	balance_shoot float64
}

func kassir1(input1 Client, input2 Product, ch chan float64) {
	if input1.balance_shoot >= input2.price {
		input1.balance_shoot = input1.balance_shoot - input2.price
	}
	ch <- input2.price
}

func kassir2(input1 Client, input2 Product, ch chan float64) {
	if input1.balance_shoot >= input2.price {
		input1.balance_shoot = input1.balance_shoot - input2.price
	}
	ch <- input2.price
}

func kassir3(input1 Client, input2 Product, ch chan float64) {
	if input1.balance_shoot >= input2.price {
		input1.balance_shoot = input1.balance_shoot - input2.price
	}
	ch <- input2.price
}

func kassir4(input1 Client, input2 Product, ch chan float64) {
	if input1.balance_shoot >= input2.price {
		input1.balance_shoot = input1.balance_shoot - input2.price
	}
	ch <- input2.price
}

func main() {
	user1 := Client{
		"Farxod",
		3000000,
	}

	user2 := Client{
		"Erkin",
		6000000,
	}

	user3 := Client{
		"Rustam",
		7000000,
	}

	user4 := Client{
		"Feruz",
		8000000,
	}

	Product1 := Product{
		"banan",
		15000,
	}

	Product2 := Product{
		"mandarin",
		17000,
	}

	Product3 := Product{
		"ananas",
		25000,
	}

	Product4 := Product{
		"apelsin",
		22000,
	}

	tunnel1 := make(chan float64)

	tunnel2 := make(chan float64)

	tunnel3 := make(chan float64)

	tunnel4 := make(chan float64)

	go kassir1(user1, Product1, tunnel1)

	go kassir2(user2, Product2, tunnel2)

	go kassir3(user3, Product3, tunnel3)

	go kassir4(user4, Product4, tunnel4)

	x := <-tunnel1

	y := <-tunnel2

	z := <-tunnel3

	k := <-tunnel4

	fmt.Printf("%v : %v\n %v : %v\n %v : %v\n %v : %v \n", user1.name, x, user2.name, y, user3.name, z, user4.name, k)

}

// 4 ta kassir

// user

// product
// price

// user1 -> kassir1 -> summa
// user2 -> kassir2 -> summa
// user3 -> kassir3 -> summa
// user4 -> kassir4 -> summa

// Farxod   : 23000
// Jahongir : 24000
// Samandar : 82000
// Moxirbek : 40000
