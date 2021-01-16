package main

import (
	"fmt"

	"gopkg.in/kyokomi/emoji.v1"
)

func main() {
	fmt.Println("Calling the function: main")
	anotherFunction()

	thirdPartyFunction()
}

func anotherFunction() {
	fmt.Println("Calling the function: anotherFunction")
}

func thirdPartyFunction() {
	fmt.Println("Hello World Emoji. :tada:")
	emoji.Println("Do you like :coffee:? Of course you like!")
}
