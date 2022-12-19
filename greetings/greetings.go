package greetings

import "fmt"

// Hello returns a greeting for the named person
// In Go, a function that start with a capital letter can be called by a function in the same package. Also known as exported name
func Hello(name string) string {
	//return a greeting that embeds the name in a message
	// := declared and initializes a variable
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
