package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	//function named "Hello", "name" parameter has string type, and function also expect return type is string as well
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
	// ":=" operator use to declare and initialize a variable in 1 line 
    return message
}