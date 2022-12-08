package main

import "fmt"

/*func main() {
	var a, b int = 4, 5

	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b)

}
*/

/*2-masala
func main() {
	var a int = 37
	fmt.Println((a%10)*10 + int(a/10))

}*/

/*3-masala

func main() {
	var a, b, c int = 3, 4, 11
	if a+b > c && a+c > b && b+c > a {
		fmt.Println("Ha")
	} else {
		fmt.Println("Yo`q")
	}

}
*/
//4-masala
func main() {
	var a, b, c int = 7, 8, 6
	if a > b && a > c {
		fmt.Println("a eng kATTASI")
	} else if b > c {
		fmt.Println("b eng kattasi")
	} else {
		fmt.Println("C eng kattasi")
	}
}
