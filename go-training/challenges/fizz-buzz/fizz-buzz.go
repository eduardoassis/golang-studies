package main

import (
	"fmt"
)

func main() {

	for x := 1; x <= 30; x++ {
		switch {
		case x%3 == 0 && x%5 == 0:
			fmt.Println("fizz-buzz")
		case x%5 == 0:
			fmt.Println("buzz")
		case x%3 == 0:
			fmt.Println("fizz")
		default:
			fmt.Printf("%v \n", x)
		}

	}

}
