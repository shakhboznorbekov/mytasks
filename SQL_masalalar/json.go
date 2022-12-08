// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )

// type Product struct {
// 	Id          int     `json:"id"`
// 	Title       string  `json:"title"`
// 	Price       float64 `json:"price"`
// 	Description string  `json:"description"`
// 	Category    string  `json:"category"`
// 	Image       string  `json:"image"`
// 	Rating      Rating  `json:"rating"`
// }

// type Rating struct {
// 	Rate  float64 `json:"rate"`
// 	Count int     `json:"count"`
// }

// func main() {

// 	var products []Product

// 	body, err := os.ReadFile("./product.json")
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = json.Unmarshal(body, &products)
// 	if err != nil {
// 		panic(err)
// 	}

// 	for _, product := range products {
// 		// if product.Rating.Count > 100 {
// 		// 	fmt.Println(product)
// 		// }

// 		if product.Rating.Rate > 5 || product.Price > 500 {
// 			fmt.Println(product.Rating.Rate, product.Price)
// 		}

// 		// if strings.Contains(product.Title, "Gold") == true {
// 		// 	fmt.Println(product.Title)
// 		// }

// 	}
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )

// type User []struct {
// 	Address  Address `json:"address"`
// 	ID       int     `json:"id"`
// 	Email    string  `json:"email"`
// 	Username string  `json:"username"`
// 	Password string  `json:"password"`
// 	Name     string  `json:"name"`
// 	Phone    string  `json:"phone"`
// 	V        int     `json:"__v"`
// }
// type Address struct {
// 	Geolocation Geolocation `json:"geolocation"`
// 	City        string      `json:"city"`
// 	Street      string      `json:"stret"`
// 	Number      int         `json:"number"`
// 	Zipcode     string      `json:zipcode"`
// }

// type Geolocation struct {
// 	Lat  string `json:"lat"`
// 	Long string `json:"long"`
// }

// func main() {
// 	client := http.Client{}
// 	rec, err := http.NewRequest("Get", "http://fakestoreapi.com/users", nil)
// 	if err != nil {
// 		panic(err)
// 	}

// 	resp, err := client.Do(rec)
// 	if err != nil {
// 		panic(err)
// 	}

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		panic(err)
// 	}
// 	var users []User
// 	err = json.Unmarshal(body, &users)

// 	for _, user := range users {
// 		fmt.Println(user)
// 	}
// }

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Application struct {
	Orders []Orders `json:"orders"`
	Count  int      `json:"count"`
}
type Orders struct {
	ID             string    `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	ThirdName      string    `json:"third_name"`
	PhoneNumber    string    `json:"phone_number"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	Speciality     string    `json:"speciality"`
	EducationType  string    `json:"education_type"`
	PassportSeries string    `json:"passport_series"`
	PassportNumber string    `json:"passport_number"`
	PassportPinfl  string    `json:"passport_pinfl"`
	CertificateURL string    `json:"certificate_url"`
	DiplomURL      string    `json:"diplom_url"`
	ImageURL       string    `json:"image_url"`
	CreatedAt      string    `json:"created_at"`
	UpdatedAt      string    `json:"updated_at"`
}

// func (o Application)Search(key, value string) interface{}{
// 	for _, application := range o.Orders {
// 		if strings.Contains(application.FirstName, "Palema") == true {
// 			return application
// 		}
// 	}

// }

func main() {
	var Applications Application

	body, err := os.ReadFile("./application.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &Applications)
	if err != nil {
		panic(err)
	}

	// struct hosil qilish
	for _, applic := range Applications.Orders {
		fmt.Println(applic)
	}

	// 	name boyicha aniqlash
	for _, application := range Applications.Orders {
		if strings.Contains(application.FirstName, "Palema") == true {
			fmt.Println(application)
		}
	}

	// strings.Split("1999-01-11T00:00:00Z","1999-01-11")

	// if Orders.DateOfBirth > "1999-01-11T00:00:00Z" || Orders.DateOfBirth < "2000-05-18T00:00:00Z" {
	// 	fmt.Println(Orders.DateOfBirth)
	// }

}

// }
