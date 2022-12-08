package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	encoder := json.NewEncoder(os.Stdout)

	encoder.Encode([]int{1, 2, 3})

	fmt.Fprint(os.Stdout, "Hello world")
}
