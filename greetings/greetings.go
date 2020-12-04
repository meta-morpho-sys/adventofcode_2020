package greetings

import (
	"fmt"
	"errors"
)

// Hello returns a greeting to a named person
func Hello(name string) string {
    // If no name was given, return an error with a message.
	if name == "" {
		return '', errors.New("name is empty")
	}

    // Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}