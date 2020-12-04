package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	//	Set properties of the predefined logger, including 
	//	the log entry prefix and a flag to disable printing
	//	the time, source file, and line number. 
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a message
	message, err := greetings.Hello("")
	//	If an error is returned, print it out to the console 
	//	and exit the program
	if err != nil {
		errors.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	
	fmt.Println(message)
}