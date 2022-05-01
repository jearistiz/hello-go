package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"rsc.io/quote"

	"example/greetings"
)

const inputName bool = false

func main() {
	simpleMain()
	errorHandlingMain()
}

func simpleMain() {
	log.SetPrefix("simpleMain: ")
	logStartExec()

	message := greetings.Hello("Jose David")
	fmt.Println(message)
	fmt.Println(quote.Go())

	logEndExec()
}

func errorHandlingMain() {
	log.SetPrefix("errorHandlingMain: ")
	logStartExec()

	var fullName string
	var err error

	if inputName {
		// Scan name from terminal
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">>> Type your name, then press Enter: ")
		fullName, err = reader.ReadString('\n')
		fullName = fullName[:len(fullName)-1]
		// Handling error from ReadString function
		if err != nil {
			logFatalEndWithError(err)
		}
	} else {
		fullName = "Juan Esteban"
	}

	// Call HelloWithError funcition using user input name
	message, err := greetings.HelloWithError(fullName)
	// Handling error from HelloWithError function
	if err != nil {
		logFatalEndWithError(err)
	}

	fmt.Println(message)

	logEndExec()
}

func logStartExec() {
	fmt.Println()
	log.Println("--- Start execution ---")
}

func logEndExec() {
	log.Println("--- End execution ---")
}

func logFatalEndWithError(err error) {
	msg := fmt.Sprintf("--- End execution with errors: %v---", err)
	log.Fatal(msg)
}
