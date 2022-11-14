package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var formats = []string {
	"Hey, %v!",
	"Nice to see you, %v!",
	"Well look who showed up. It's %v!",
}

func Hello(input string) (string, error) { // The function "Hello" is an "Exported Name" (accessible outside this package)
	if input == "" {
		return "", errors.New("Parameter 'input' cannot be empty.")
	}

	 message := fmt.Sprintf(getRandomFormatString(), input)
	 return message, nil
}

// This function is run auto-magically after "all the variable declarations in the package have evaluated 
// their initializers, and those are evaluated only after all the imported packages have been initialized."
// See: https://go.dev/doc/effective_go#init
func init() { 
	rand.Seed(time.Now().UnixNano())
}

func getRandomFormatString() string {
	return formats[rand.Intn(len(formats))]
}
