package main

import (
	"fmt"

	"rsc.io/quote"

	"example.com/greetings"

	"log"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings : ")
	log.SetFlags(0)
	fmt.Println("Hello, World!")
	message, err := greetings.Hello("Prakhar Srivastava")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
	fmt.Println(quote.Go())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)

	message1, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message1)
}
