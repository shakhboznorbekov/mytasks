// package main

// import (
// 	"fmt"
// 	"time"
// )

// func TashkentAirways(ch chan string) {

// 	ch <- "Message from Tashkent"
// }

// func IstanbulAirways(ch chan string) {

// 	time.Sleep(10 * time.Millisecond)
// 	ch <- "Message from Istanbul"
// }

// func main() {

// 	tashkent_airways := make(chan string)
// 	istanbul_airways := make(chan string)

// 	go TashkentAirways(tashkent_airways)
// 	go IstanbulAirways(istanbul_airways)

// 	select {
// 	case v := <-tashkent_airways:
// 		fmt.Println(v)
// 	case v := <-istanbul_airways:
// 		fmt.Println(v)
// 	}
// }
