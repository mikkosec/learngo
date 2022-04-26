package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	// Error handling
	if name == "" {
		return "", errors.New("empty name")
	}

	// Create message using a random format.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates each of the named people with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	// Loop through the received slice of names, calling the hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with the name
		messages[name] = message
	}
	return messages, nil
}

// init sets initial values for variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages.
func randomFormat() string {
	// Slice of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v!",
	}
	// return a randombly selected message format by specifying a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
