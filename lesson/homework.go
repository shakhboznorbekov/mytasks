package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 1-masala
// func main() {
// 	var musbat = []int{}
// 	var manfiy = []int{}
// 	i := 0
// 	for i < 500 {
// 		son := rand.Intn(2000) - 1000
// 		if son > 0 {
// 			musbat = append(musbat, son)
// 		} else {
// 			manfiy = append(manfiy, son)
// 		}
// 		i++
// 	}
// 	fmt.Println(musbat)
// 	fmt.Println(manfiy)
// }

//2-masala

// func main() {
// 	var text string = "Uyga vazifaRandom sonlar elon qiling. -1000 va 1000 gachahuni ichidagi minus sonlarni alohida slice ga joylashtiringva musbatlarni ham alohida slice ga joylashtiring.Va ekranga tartiblangan holatida chiqazingnums[500] = rand(-1000 dan 1000 gacha)2. Text qabul qiling va shuni teskari holatida qaytaruvchi function elong qilingva ekranga chiqazing3slice e'lon qiling va unga random sonlar 50 soling	Shu sonlar ichidan faqat tublarini chiqaruvchi funksiya tuzing.4. slice e'lon qiling va unga random sonlar soling.Shu sonlarng eng kattasini qaytaruvchi funksiya tuzing5. slice e'lon qiling va unga random sonlar soling.Shu sonlar ichidan 2-tub sonni qaytaruvchi funksiya tuzing."
// 	i := 500
// 	for i >= 0 {
// 		fmt.Printf("%s", string(text[i]))
// 		i--
// 	}
// }

// //3-masala

// func tub(a []int) []int {
// 	var tublar []int
// 	for _, v := range a {
// 		check := true
// 		for i := 2; i < v; i++ {
// 			if v%i == 0 {
// 				check = false
// 				break
// 			}
// 		}
// 		if check && v != 0 {
// 			tublar = append(tublar, v)
// 		}
// 	}
// 	return tublar
// }

// func main() {
// 	var nums []int
// 	i := 0
// 	for i < 50 {
// 		nums = append(nums, rand.Intn(50))
// 		i++
// 	}
// 	fmt.Println(tub(nums))
// 	fmt.Println(nums)

// }

// // 4-masala
// func Max(a []int) int {
// 	max := a[0]
// 	for _, v := range a {
// 		if max < v {
// 			max = v
// 		}
// 	}
// 	return max
// }

// func main() {
// 	var nums []int
// 	for i := 0; i < 50; i++ {
// 		nums = append(nums, rand.Intn(50))
// 	}
// 	fmt.Println(Max(nums))
// 	fmt.Println(nums)
// }

//5-masala

func Tub(a []int) int {
	var sum []int
	for _, v := range a {
		check := true
		for i := 2; i < v; i++ {
			if v%i == 0 {
				check = false
				break
			}
		}
		if check && v != 0 {
			sum = append(sum, v)
		}
	}
	sort.Ints(sum)
	return sum[1]
}

func main() {
	var nums []int
	for i := 0; i < 50; i++ {
		nums = append(nums, rand.Intn(50))
	}
	fmt.Println(Tub(nums))
	fmt.Println(nums)
}
