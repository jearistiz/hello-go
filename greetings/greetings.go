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
	"Greetings, %v. It's a pleasuere. We are ON THE GO!",
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

func Hello(name string) string {

	message := fmt.Sprintf(defaultHelloMsg, name)

	// The above LOC is equivalent to the following two commented LsOC
	// var message string
	// message = fmt.Sprintf("Hi, %v. Welcome!", name)

	return message

}

func HelloWithError(name string) (string, error) {
	err := validateName(name)
	if err == nil {
		return Hello(name), err
	} else {
		return "", err
	}
}

func RandomHello(name string) (string, error) {
	err := validateName(name)
	if err == nil {
		randomHelloMsg := randomFormattedHelloMsg()
		return fmt.Sprintf(randomHelloMsg, name), err
	} else {
		return "", err
	}
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
