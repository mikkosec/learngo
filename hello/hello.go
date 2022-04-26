package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined logger.
	log.SetPrefix("greetings: ")
	log.SetFlags(0) // Flag to disable printing

	// Request a greeting message
	message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
