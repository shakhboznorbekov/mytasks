package main

import (
	"fmt"
)

type Identification struct {
	firstName       string
	lastName        string
	birthday        string
	fullName        string
	age             int
	groupcoursename string
	groupduration   string
}

func main() {
	var user Identification
	user = Identification{
		firstName:       "Shaxboz",
		lastName:        "Norbekov",
		birthday:        "22.11.1998",
		fullName:        "Norbekov Shaxboz",
		age:             23,
		groupcoursename: "go Beckend",
		groupduration:   "3 oy",
	}

	fmt.Println(user.firstName)
	fmt.Println(user.lastName)
	fmt.Println(user.birthday)
	fmt.Println(user.fullName)
	fmt.Println(user.age)
	fmt.Println(user.groupcoursename)
	fmt.Println(user.groupduration)
}
