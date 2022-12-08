package main

import "fmt"

type AcademyRoom struct {
	Door       string
	WhiteBoard string
	Table      uint8
	Chair      uint8
}

func main() {
	var hello []AcademyRoom

	for i := 0; i < 10; i++ {
		var uacademy AcademyRoom = AcademyRoom{
			Door:       "Room 1",
			WhiteBoard: "White Board",
			Table:      7,
			Chair:      13,
		}
		hello = append(hello, uacademy)
	}
	fmt.Println(hello[1])
}
