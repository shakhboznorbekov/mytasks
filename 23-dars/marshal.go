//godan json ga otadi

package main

import (
	"encoding/json"
	// "fmt"
	"os"
)

type user struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
	Age         int    `jon:"age"`
}

func main() {
	// var u user = user{
	// 	FirstName:   "Shaxboz",
	// 	LastName:    "Norbekov",
	// 	PhoneNumber: "+998993074358",
	// 	Age:         24,
	// }

	// body, _ := json.Marshal(u)

	// fmt.Println(string(body))

	// a, _ := json.Marshal(123)
	// b, _ := json.Marshal(true)
	// c, _ := json.Marshal("text")
	// d, _ := json.Marshal(32.3242)
	// e, _ := json.Marshal('A')
	// f, _ := json.Marshal(nil)
	// fmt.Println(string(a), string(b), string(c), string(d), string(e), string(f))

	data, err := json.Marshal([]int{1, 2, 3})
	// fmt.Println(string(data))

	if err != nil {
		panic(err)
	}

	os.Stdout.Write((data))
}
