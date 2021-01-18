package main

import (
	"fmt"
)

func main() {

	book := "The colour of magic"

	fmt.Println(len(book)) // use lent function to know how many byte has a string

	// Use '[]' to access individual bytes of a string
	// uint8 is byte in go
	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])

	// strings in go are immutable
	//book[0] = 116

	// Slice (start, end), 0 based, half empty range
	fmt.Println(book[1:11])

	// Slice no end
	fmt.Println(book[4:])

	// Slice no start
	fmt.Println(book[:4])

	// Contatenete strings with + operator
	fmt.Println("t" + book[1:])

	// Strings in go are Unicode
	fmt.Println("It was ½ price!")

	// Multi line
	poem := `
	The road goes ever on
	Down from the door where it began`
	fmt.Println(poem)
}
