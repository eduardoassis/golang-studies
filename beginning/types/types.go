package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(3))
	fmt.Println(reflect.TypeOf(3.14))
	fmt.Println(reflect.TypeOf("Just a String"))
	fmt.Println(reflect.TypeOf(true))
}
