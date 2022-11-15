package main

import (
	"fmt"
	"log"

	"github.com/Phogheus/GoLearnGo/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Awesome Name")

	if err != nil {
		log.Fatal(err) // Calls Print and then exits application with error code 1
	}

	fmt.Println(message) // Prints: "Hey, Awesome Name!"
}
