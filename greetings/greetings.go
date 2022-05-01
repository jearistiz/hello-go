package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const defaultHelloMsg string = "Hi, %v. Welcome to ON THE GO!"

var helloFormattedMsgs = [...]string{
	defaultHelloMsg,
	"Hello, %v. We are ON THE GO! Are you?",
	"Greetings, %v. Do you wanna be ON THE GO?",
	"Great to see you, %v! Let's get ON THE GO!",
}

// From go website: Go executes init functions automatically at program startup,
// after global variables have been initialized. For more about init functions,
// see [Effective Go](https://go.dev/doc/effective_go#init).
//
// In this case, init() sets the seed for the rand module using the current time.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Hello returns a greeting for the named person.
func Hello(name string) string {

	message := fmt.Sprintf(defaultHelloMsg, name)

	// The above LOC is equivalent to the following two commented LsOC
	// var message string
	// message = fmt.Sprintf("Hi, %v. Welcome!", name)

	return message

}

// Hello returns a greeting for the named person and an error if
// name is empty.
func HelloWithError(name string) (string, error) {
	err := validateName(name)
	if err == nil {
		return Hello(name), err
	} else {
		return "", err
	}
}

// Hello returns a random greeting for the named person.
func RandomHello(name string) (string, error) {
	err := validateName(name)
	if err == nil {
		randomHelloMsg := randomFormattedHelloMsg()
		return fmt.Sprintf(randomHelloMsg, name), err
	} else {
		return "", err
	}
}

// RandomHellos returns a map that associates each of the named people
// with a greeting message.
func RandomHellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := RandomHello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func randomFormattedHelloMsg() string {
	return helloFormattedMsgs[rand.Intn(len(helloFormattedMsgs))]
}

func validateName(name string) error {
	if name == "" {
		return errors.New("empty name")
	}
	return nil
}
