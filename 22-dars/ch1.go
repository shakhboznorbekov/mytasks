// package main

// import (
// 	"sync"
// )

// func A(wg *sync.WaitGroup) {
// 	wg.Done()
// }

// func B(wg *sync.WaitGroup) {
// 	wg.Done()
// }

// func C(wg *sync.WaitGroup) {
// 	wg.Done()
// }

// func main() {

// 	var wg sync.WaitGroup

// 	go A(&wg)
// 	wg.Add(1)

// 	go B(&wg)
// 	wg.Add(1)

// 	go C(&wg)
// 	wg.Add(1)

// 	wg.Wait()
// }
