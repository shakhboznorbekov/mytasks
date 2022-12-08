package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Numbers(w http.ResponseWriter, r *http.Request) {

	encoder := json.NewEncoder(w)

	w.Header().Set("Content-Type", "application/json")

	encoder.Encode([]int{1, 2, 3, 4, 5})
}

func Names(w http.ResponseWriter, r *http.Request) {

	encoder := json.NewEncoder(w)

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-Type", "application/json")

	encoder.Encode([]string{"Apple", "Google", "Meta", "Yandex"})
}
func main() {

	fmt.Println("Listening server...")

	// http.HandleFunc("/numbers", func(w http.ResponseWriter, r *http.Request) {
	// 	encoder := json.NewEncoder(w)

	// 	encoder.Encode([]int{})
	// })

	http.HandleFunc("/numbers", Numbers)

	http.HandleFunc("/names", Names)

	http.ListenAndServe(":4000", nil)
}
