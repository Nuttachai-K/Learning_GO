package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Get a greeting message and print it.
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//A slice of name
	names := []string{"Gladys", "Samantha", "Darrin"}
	//Request greeting message for each name.
	messages, err := greetings.Hellos(names)

	//message, err := greetings.Hello("Gladys")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned map of message
	// to the console.
	fmt.Println(messages)

}
