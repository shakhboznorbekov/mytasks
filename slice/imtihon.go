package main

import (
	"fmt"
)

//1-masala
// func point(x, y, z int) (float64, float64, float64) {
// 	return float64(x) + 0.5, float64(y) + 0.5, float64(z) + 0.5
// }

//	func main() {
//		x, y, z := point(2, 3, 4)
//		fmt.Println(x)
//		fmt.Println(y)
//		fmt.Println(z)
//	}
//
// 2-masala

func request(code int, message string) (int, string) {
	return 200, "OK"
}
func main() {

	code, message := request("/users")
	fmt.Println(code)
	fmt.Println(message)
}

//3-masala

// type User struct {
// 	Name    string
// 	Surname string
// }

// func (u User) fullName() string {
// 	return u.Name + " " + u.Surname
// }

// func main() {
// 	user := User{"Asadbek", "Ergashev"}
// 	fmt.Println(user.fullName())

// }

// 4-masala
// func Sum(num ...int) int {
// 	var result int
// 	for _, val := range num {
// 		result += val
// 	}
// 	return result
// }

// func main() {
// 	sum := Sum(1, 2, 3, 4, 5, 21)

// 	fmt.Println(sum)
// }

// 5-masala
// type i64 int64
// type text string
// type u8 uint8

// func aliasing(a int64, b string, c uint8) (i64, text, u8) {
// 	return 0, "X", 255
// }
// func main() {
// 	a, b, c := aliasing(0, "X", 255)
// 	fmt.Println(a, b, c)
// }

// //6-masala
// func super(a int) func(){}{
// 	return
// }

// func main(){

// 	a:=super(1)(2)(3)
// 	b:=super(4)(5)(6)
// 	c:=super(7){8}(9)

// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// }
