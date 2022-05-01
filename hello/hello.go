package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"rsc.io/quote"

	"example/greetings"
)

func main() {
	simpleMain()
	errorHandlingMain()
}

func simpleMain() {
	fmt.Println("\n--- Start simpleMain exec ---")

	message := greetings.Hello("Jose David")
	fmt.Println(message)
	fmt.Println(quote.Go())

	fmt.Println("--- End simpleMain exec ---")
}

func errorHandlingMain() {

	fmt.Println("\n--- Start errorHandlingMain exec ---")
	// Sets the Logger prefix
	log.SetPrefix("greetings: ")
	// Avoids printing time, source file and line number in logs
	// log.SetFlags(0)

	// Scan name from terminal
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Type your name, then press Enter: ")
	fullName, err := reader.ReadString('\n')

	// Handling error from ReadString function
	if err != nil {
		msg := fmt.Sprintf("\n--- End errorHandlingMain exec with errors: %v---", err)
		log.Fatal(msg)
	}

	// Call HelloWithError funcition using user input name
	message, err := greetings.HelloWithError(fullName[:len(fullName)-1])
	// Handling error from HelloWithError function
	if err != nil {
		msg := fmt.Sprintf("\n--- End errorHandlingMain exec with errors: %v---", err)
		log.Fatal(msg)
	}

	fmt.Println(message)

	fmt.Println("--- End errorHandlingMain exec ---")
}
