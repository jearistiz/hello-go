package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	// The above loc is equivalent to the following two commented loc
	// var message string
	// message = fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
