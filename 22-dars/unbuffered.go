// package main

// import (
// 	"fmt"
// )

// func main() {

// 	// unbuffered
// 	ch := make(chan int)

// 	go func() {
// 		ch <- 1
// 	}()

// 	<-ch

// 	fmt.Println("OK")
// }
