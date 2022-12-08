// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// type user struct {
// 	Id   int
// 	Name string
// }

// func infoPage(w http.ResponseWriter, r *http.Request) {

// 	if r.Method == "GET" {

// 		u := user{
// 			Id:   1,
// 			Name: "Moxirbek",
// 		}

// 		fmt.Fprintf(w, fmt.Sprintf("User Id: %v\nUser name: %v\n",
// 			u.Id,
// 			u.Name,
// 		))
// 	}

// 	// fmt.Fprintf(w, "Info page")
// 	fmt.Println("Endpoint /info page")
// }

// func main() {

// 	http.HandleFunc("/info", infoPage)

// 	fmt.Println("Listening...")
// 	http.ListenAndServe(":8080", nil)
// }
