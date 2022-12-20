package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person
// In Go, a function that start with a capital letter can be called by a function in the same package. Also known as exported name
func Hello(name string) (string, error) {
	//if no name was given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}

	//Create a message using a random format
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// init sets initial values for variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of  a set of greeting messages. The returned message is selected at random
func randomFormat() string {
	// A slice of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Greate to see you, %v",
		"Hail, %v well met!",
	}

	//Return a randomly selected format by specifying a random index for the slice of formats
	return formats[rand.Intn(len(formats))]
}
