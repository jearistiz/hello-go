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
	randomHelloMain()
	randomHellosMain()
}

const useInputName bool = false
const defaultName string = "Juan Esteban"
const msgInputName string = ">>> Type your name, then press Enter: "

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

	if useInputName {
		fullName = inputFromTerminal(msgInputName)
	} else {
		fullName = defaultName
	}

	// Common error handling in Go: the function returns an error and the
	// client checks for the error
	message, err := greetings.HelloWithError(fullName)
	if err != nil {
		logFatalEndWithError(err)
	}

	fmt.Println(message)

	logEndExec()
}

func randomHelloMain() {
	log.SetPrefix("randomHelloMain: ")
	logStartExec()

	var fullName string

	if useInputName {
		fullName = inputFromTerminal(msgInputName)
	} else {
		fullName = defaultName
	}

	message, err := greetings.RandomHello(fullName)
	if err != nil {
		logFatalEndWithError(err)
	}

	fmt.Println(message)

	logEndExec()
}

func randomHellosMain() {
	log.SetPrefix("randomHellosMain: ")
	logStartExec()

	var fullNames []string = []string{
		"Luisa",
		"Juan",
		"Gil",
		"Venus",
		"José",
		"Emma",
	}

	messages, err := greetings.RandomHellos(fullNames)

	if err != nil {
		logFatalEndWithError(err)
	}

	for name, message := range messages {
		fmt.Printf("%v: %v\n", name, message)
	}

	logEndExec()
}

func inputFromTerminal(inputMsg string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(inputMsg)
	fullName, err := reader.ReadString('\n')
	if err != nil {
		logFatalEndWithError(err)
	}
	return fullName[:len(fullName)-1]
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
