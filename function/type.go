package main

import (
	"fmt"
	"reflect"
)

func greet(n interface{}) {
	fmt.Println(reflect.TypeOf(n).Name() == "int")

}

func main() {
	greet(2)
	greet("hello")
}
