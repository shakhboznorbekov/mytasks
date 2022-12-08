//jsondan go ga otish uchun

package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
	Age         int    `jon:"age"`
	IsActive    bool   `json:"isActive"`
	books       []int  `json:"books"`
}

var json_content = `
{
	"firstName": "Shaxboz",
	"lastName": "Norbekov",
	"age": 19,
	"isactive": true,
	"books":[]
}
`

func main() {
	var user User

	err := json.Unmarshal([]byte(json_content), &user)
	if err != nil {
		panic(err)
	}

	fmt.Println(user.FirstName)
}
