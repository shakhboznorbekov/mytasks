package main

import (
	// "encoding/json"
	"fmt"
)

// func sonlar(a int) string {
// 	birlik := map[int]string{
// 		1:  "Yanvar",
// 		2:  "Fevral",
// 		3:  "Mart",
// 		4:  "Aprel",
// 		5:  "May",
// 		6:  "Iyun",
// 		7:  "Iyul",
// 		8:  "Avgust",
// 		9:  "Sentabr",
// 		10: "Oktabr",
// 		11: "Noyabr",
// 		12: "Dekabr",
// 	}
// 	if a == 0 {
// 		panic(0)
// 	} else if a > 12 {
// 		return ""
// 	} else {
// 		return birlik[a]
// 	}
// }

// func main() {
// 	var a int
// 	for {
// 		fmt.Scanf("%d", &a)
// 		if a == 0 {
// 			break
// 		}
// 		fmt.Println(sonlar(a))
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"reflect"
// 	"strconv"
// )

// func main() {
// 	num1 := "6"
// 	num2 := "3"
// 	c := 0

// 	i, err := strconv.Atoi(num1)
// 	if err != nil {
// 		panic(err)
// 	}
// 	z, err := strconv.Atoi(num2)
// 	if err != nil {
// 		panic(err)
// 	}
// 	c = i + z
// 	str := strconv.Itoa(c)
// 	fmt.Println(str, reflect.TypeOf(str))

// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	var son string = "12@#$%"
// 	result := strings.Trim(son, "$%")
// 	fmt.Println(result)
// }

func week(a int) {
	days := []string{"M", "Th", "W", "T", "F", "S", "Su"}
	fmt.Println(days[a-1])
}

func mark(a int) {
	if a < 1 || a > 5 {
		fmt.Println("xato")
	} else {
		marks := []string{"Yomon", "Qoniqarsiz", "Qoniqarli", "Yaxshi", "A'lo"}
		fmt.Println(marks[a-1])
	}
}

func month(a int) {
	season := []string{"Qish", "Bahor", "Yoz", "Kuz"}
	if a == 1 || a == 2 || a == 12 {
		fmt.Printf("%d chi oy,'%v'\n", a, season[0])
	} else if a == 3 || a == 4 || a == 5 {
		fmt.Printf("%d chi oy,'%v'\n", a, season[1])
	} else if a == 6 || a == 7 || a == 8 {
		fmt.Printf("%d chi oy,'%v'\n", a, season[2])
	} else if a == 9 || a == 10 || a == 11 {
		fmt.Printf("%d chi oy,'%v'\n", a, season[3])
	} else {
		fmt.Println("12ta oy mavjud!!!")
	}

}

func NumbersMonths(num int) {
	switch num {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println("31 kun")
	case 4, 6, 9, 11:
		fmt.Println("30 kun")
	case 2:
		fmt.Println("29/28 kun")
	default:
		fmt.Println("Sorry")
	}
}

func calc(a, b float64, amal int) {
	switch amal {
	case 1:
		fmt.Println(a + b)
	case 2:
		fmt.Println(a - b)
	case 3:
		fmt.Println(a / b)
	case 4:
		fmt.Println(a * b)
	default:
		fmt.Println("Faqat 4ta amalni bilaman 'sorry'")
	}
}

func main() {
	// week(5)
	// mark(10)
	// month(13)
	// NumbersMonths(100)
	calc(2.1, 21.2, 1)

}
