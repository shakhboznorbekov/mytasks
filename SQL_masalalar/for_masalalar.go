package main

import "fmt"

// // 1-masala

// //    1
// //    12          korinishidagi misol
// //    123
// //    1234

// // func tringleNum(n int) {
// // 	var sum int = 0
// // 	for i := 1; i <= n; i++ {
// // 		sum = sum*10 + i
// // 		fmt.Println(sum)
// // 	}
// // }

// // func main() {
// // 	var son int = 4
// // 	tringleNum(son)
// // }

// //2-masala

// //                                                *
// //                                               * *
// //                                              * * *        <--  like this
// //                                             * * * *
// //                                            * * * * *
// //                                           * * * * * *
// //                                          * * * * * * *
// //                                         * * * * * * * *
// //                                        * * * * * * * * *
// //                                       * * * * * * * * * *
// // func PiramidNum(n int) {
// // 	space := n - 1
// // 	star := 1

// // 	for i := 1; i <= n; i++ {

// // 		for j := 1; j <= space; j++ {
// // 			fmt.Printf(" ")
// // 		}

// // 		for k := 1; k <= star; k++ {
// // 			fmt.Printf("* ")
// // 		}
// // 		space--
// // 		star++
// // 		fmt.Println()
// // 	}
// // }

// // func main() {
// // 	var son int = 10
// // 	PiramidNum(son)
// // }

// //3-masala

// // func main() {
// // 	var sum int
// // 	for i := 100; i < 200; i++ {

// // 		if i%9 == 0 {
// // 			fmt.Println(i)
// // 			sum += i
// // 		}
// // 	}
// // 	fmt.Println(sum)
// // }

func son(a int) {
	for i := a; i > 0; i-- {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func son2(a int) {
	for i := a; i > 0; i-- {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}

func tub() {
	c := 0
	for i := 10; i < 100; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				c++
				if c < 1 {
					fmt.Println(i)
				}
				c = 0
			}
		}
	}
}

func takrorlanuvchi() {
	a, b, c := 0, 0, 0
	for i := 100; i < 1000; i++ {
		k := i
		a = i % 10
		i = i / 10
		b = i % 10
		i = i / 10
		c = i
		if a == b || b == c || a == c {
			fmt.Println(k)

		}
		i = k
	}
}

func teskari(a int) {
	for a >= 1 {
		fmt.Printf("%d", a%10)
		a = a / 10
	}
	fmt.Println()
}

//	func teskari(a int) {
//		for a > 1 {
//			if a%10 != 0 {
//				fmt.Printf("%d", a%10)
//			}
//			a = a / 10
//		}
//	}

// UYGA VAZIFA!!!!
func main() {
	// son(30)       //1-masala
	// son2(30)      //2-masala
	// tub()         //3-masala
	// takrorlanuvchi() //4-masala
	// teskari(1123) //5-masala

}

// 6-masala
// func PiramidNum(n int) {
// 	space := n - 1
// 	star := 1

// 	for i := 1; i <= n; i++ {

// 		for j := 1; j <= space; j++ {
// 			fmt.Printf(" ")
// 		}

// 		for k := 1; k <= star; k++ {
// 			fmt.Printf("* ")
// 		}
// 		space--
// 		star++
// 		fmt.Println()
// 	}
// }

// func main() {
// 	var son int = 10
// 	PiramidNum(son)
// }
