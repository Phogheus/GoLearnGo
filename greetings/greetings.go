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

func MultipleHellos(names []string) (map[int]string, error) {
	messages := make(map[int]string) // Create (make) a new map (Name/Value pair)

	for i, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[i] = message
	}

	return messages, nil
}

// This function is run auto-magically after "all the variable declarations in the package have evaluated 
// their initializers, and those are evaluated only after all the imported packages have been initialized."
// See: https://go.dev/doc/effective_go#init
func init() { 
	rand.Seed(time.Now().UnixNano())
}

// The function "getRandomFormatString" is /not/ an "Exported Name" (accessible inside this package only)
func getRandomFormatString() string {
	return formats[rand.Intn(len(formats))]
}
