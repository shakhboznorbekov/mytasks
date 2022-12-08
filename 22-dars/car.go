package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var a int = 1

type Student struct {
	Name string
}

func (Student) GetSpeed() time.Duration {

	rand.Seed(time.Now().UnixNano())

	speed := rand.Intn(5000)

	return time.Duration(speed) * time.Millisecond
}

func Run(s Student, wg *sync.WaitGroup) {

	time.Sleep(s.GetSpeed())

	fmt.Printf("%d.%v inshoni tugatdi.\n", a, s.Name)
	a += 1
	wg.Done()
}

func main() {
	var a uint
	fmt.Println("Student musobaqasiga xush kelibsiz!!!\nBugun quyidagi studentlar bellashmoqda")
	time.Sleep(2 * time.Second)
	fmt.Println("Mohirbek\nSamandar\nBobur\nRavshanbek\nJahongir\nShaxboz")
	time.Sleep(2 * time.Second)
	fmt.Println(" Inshoni boshlash uchun <<1>>ni bosing...")
	fmt.Scanf("%v", &a)
	if a == 1 {
		var wg sync.WaitGroup

		talaba := []Student{
			Student{"Mohirbek"},
			Student{"Samandar"},
			Student{"Bobur"},
			Student{"Ravshanbek"},
			Student{"Jahongir"},
			Student{"Shaxboz"},
		}

		for _, student := range talaba {

			wg.Add(1)
			go Run(student, &wg)
		}

		wg.Wait()
		if a == 1 {

		}

	} else {
		fmt.Println("Afsus siz boshqa raqam kiritdingiz...")
	}

}
