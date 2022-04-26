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

	// A slice of names
	names := []string{"Josephine", "Dolores", "Cathy"}

	// Request a greeting message
	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
