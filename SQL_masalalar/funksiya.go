package main

import "fmt"

// 2.
// Input any number for square : 20
// Expected Output :

// The square of 20 is : 400.00

// func square(a int) float64 {
// 	return float64(a) * float64(a)
// }

// func main() {
// 	var a int = 30
// 	fmt.Println(square(a))
// }

//3-masala
// Input 1st number : 2
// Input 2nd number : 4
// Expected Output :

// Before swapping: n1 = 2, n2 = 4
// After swapping: n1 = 4, n2 = 2

// solution
// func swap(a, b int) (int, int) {
// 	a, b = b, a
// 	return a, b

// }

// func main() {
// 	a, b := 2, 4

// 	fmt.Println(swap(a, b))
// }

//4-masala
// Input any number : 5
// Expected Output :

//  The entered number is odd.

//solution

// func num(a int) {
// 	if a%2 == 0 && a != 0 {
// 		fmt.Printf("The entered number is even:%d\n", a)
// 	} else if a == 0 {
// 		fmt.Printf("The entered number is null:%d\n", a)
// 	} else {
// 		fmt.Printf("The entered number is odd:%d\n", a)
// 	}

// }

// func main() {
// 	var a int = 27
// 	num(a)
// }

// 5-masala
// Write a program in Go to find the sum of the series 1!/1+2!/2+3!/3+4!/4+5!/5 using the function. Go to the editor
// Expected Output :

//  The sum of the series is : 34

// solution

func sum(a int) {
	count := 1
	sum := 0
	for i := 1; i <= a; i++ {
		for j := 1; j <= i; j++ {
			count *= j
		}
		sum += (count / i)
		count = 1
	}
	fmt.Printf("The sum of the series is : %d\n", sum)
}

func main() {
	a := 5
	sum(a)
}

//6-masala           SONNI IKKILIK SANOQ SISTEMASIGA O`TKAZISH  -->REKURSIYA
// Input any decimal number : 65
// Expected Output :

//  The Binary value is : 1000001
//solution

// func binary(a int) {
// 	var c int
// 	c = a % 2
// 	if a/2 >= 1 {
// 		binary(a / 2)
// 	}
// 	fmt.Printf("%d", c)

// }

// func main() {
// 	a := 65
// 	binary(a)
// 	fmt.Println()
// }

//7-MASALA           SONNI TUB YOKI TUB EMASLIGINI ANIQLASH
// Input a positive number : 5
// Expected Output :

// The number 5 is a prime number.

// solution

// func prime(a int) {
// 	count := 0
// 	for i := 2; i < a; i++ {
// 		if a%i == 0 {
// 			count++
// 		}
// 	}
// 	if count >= 1 {
// 		fmt.Printf("The number %d is not prime number.", a)
// 	} else {
// 		fmt.Printf("The number %d is a prime number.", a)
// 	}
// }

// func main() {
// 	a := 8
// 	prime(a)
// }

//8-MASALA      ARRAYNI ICHIDAGI ENG KATTA QIYMATNI TOPISH
// Input the number of elements to be stored in the array :5
// Input 5 elements in the array :
// element - 0 : 1
// element - 1 : 2
// element - 2 : 3
// element - 3 : 4
// element - 4 : 5
// Expected Output :

// The largest element in the array is : 5

//solution

// func res(a []int) {
// 	max := a[0]
// 	for ind, val := range a {
// 		fmt.Printf("element - %d : %d\n", ind, val)
// 		if max < val {
// 			max = val
// 		}
// 	}
// 	fmt.Printf("The largest element in the array is : %d\n", max)
// }
// func main() {
// 	a := []int{100, 2, 3, 4, 5, 6, 77, 8, 9, 150, 66}
// 	res(a)
// }

//9-MASALA
// Input any number: 371
// Expected Output :

//  The 371 is an Armstrong number.
//  The 371 is not a Perfect number.

//solution

// func armstrong(a int) {
// 	xona, a2 := 0, a
// 	for a2 > 0 {
// 		a2 /= 10
// 		xona++
// 	}
// 	a2 = a
// 	son, kop := 0, 1
// 	sum := 0
// 	for a2 > 0 {
// 		son = a2 % 10
// 		for i := 0; i < xona; i++ {
// 			kop *= son

// 		}
// 		sum += kop
// 		kop = 1
// 		son = 0
// 		a2 = a2 / 10

// 	}
// 	if sum == a {
// 		fmt.Printf("The %d is an Armstrong number.", a)
// 	} else {
// 		fmt.Printf("The %d is not a Perfect number.", a)
// 	}

// }

// func main() {
// 	a := 371

// 	armstrong(a)

// }

//10-masala            boluvchilari yigindisi oziga teng son
// Input lowest search limit of perfect numbers : 1
// Input lowest search limit of perfect numbers : 100
// Expected Output :

//  The perfect numbers between 1 to 100 are :
//  6   28

//solution

// func count(a int) {
// 	c := 0
// 	for i := 1; i < a; i++ {
// 		for j := 1; j < i; j++ {
// 			if i%j == 0 {
// 				c += j
// 			}
// 		}
// 		if c == i {
// 			fmt.Println(i)
// 		}
// 		c = 0
// 	}

// }

// func main() {
// 	son := 100

// 	count(son)
// }
