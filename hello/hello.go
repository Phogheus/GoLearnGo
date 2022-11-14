package main

import (
	"fmt"
	"log"

	"github.com/Phogheus/GoLearnGo/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	doGreeting("Awesome Name") // Prints: "Hello, Awesome Name!"
	doGreeting("") // Prints: "greetings: Parameter 'input' cannot be empty." followed by "exit status 1"
}

func doGreeting(input string) {
	message, err := greetings.Hello(input)

	if err != nil {
		log.Fatal(err) // Calls Print and then exits application with error code 1
	}

	fmt.Println(message)
}
