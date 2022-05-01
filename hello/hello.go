package main

import (
	"fmt"

	"rsc.io/quote"

	"example/greetings"
)

func main() {
	message := greetings.Hello("Jose David")

	fmt.Println(message)

	fmt.Println(quote.Go())
}
