package main

import (
	"fmt"
)

func main() {
	count := 0

	for i := 1000; i <= 9999; i++ {
		for j := i; j <= 9999; j++ {
			n := i * j
			numberFormatedInString := fmt.Sprintf("%d", n)
			if numberFormatedInString[0] == numberFormatedInString[len(numberFormatedInString)-1] {
				count++
			}
		}
	}
	fmt.Printf("Even ended numbers result from multiplying two four-digits numbers is: %v\n", count)
}
