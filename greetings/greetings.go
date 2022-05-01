package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) string {

	message := fmt.Sprintf("Hi, %v. Welcome to ON THE GO!", name)

	// The above LOC is equivalent to the following two commented LsOC
	// var message string
	// message = fmt.Sprintf("Hi, %v. Welcome!", name)

	return message

}

func HelloWithError(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	return Hello(name), nil
}
