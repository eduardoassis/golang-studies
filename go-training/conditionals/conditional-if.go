// Example of If statement
package main

import (
	"fmt"
)

func main() {

	x := 10

	if x > 5 {
		fmt.Println("x is bigger")
	} else {
		fmt.Println("x is smaller")
	}

	if x > 5 && x < 15 {
		fmt.Println("x is just right")
	}
}