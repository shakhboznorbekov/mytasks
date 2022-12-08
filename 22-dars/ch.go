// package main

// import (
// 	"fmt"
// )

// func receive(ch1 chan int) {

// 	for i := 1; i < 10; i++ {
// 		ch1 <- i
// 	}

// 	close(ch1)
// }

// func main() {

// 	ch := make(chan int)

// 	go receive(ch)

// 	for n := range ch {
// 		fmt.Println(n)
// 	}
// }
