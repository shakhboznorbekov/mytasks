// package main

// import (
// 	"fmt"
// 	"time"
// 	"sync"
// 	"math/rand"
// )

// type Car struct {
// 	Name string
// }

// func (Car) GetSpeed() time.Duration {

// 	rand.Seed(time.Now().UnixNano())

// 	speed := rand.Intn(5000)

// 	return time.Duration(speed) * time.Millisecond
// }

// func Run(c Car, wg *sync.WaitGroup) {

// 	time.Sleep(c.GetSpeed())

// 	fmt.Println(c.Name + " is finished.")

// 	wg.Done()
// }

// func main() {

// 	var wg sync.WaitGroup

// 	cars := []Car {
// 		Car { "BMW X 6"},
// 		Car { "Ferrari"},
// 		Car { "Matiz Best"},
// 		Car { "Yangi Damas"},
// 		Car { "Jiguli 06"},
// 	}

// 	for _, car := range cars {

// 		wg.Add(1)
// 		go Run(car, &wg)
// 	}

// 	wg.Wait()
// }