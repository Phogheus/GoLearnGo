package greetings

import (
	"fmt"
	"errors"
)

func Hello(input string) (string, error) { // The function "Hello" is an "Exported Name" (accessible outside this package)
	if input == "" {
		return "", errors.New("Parameter 'input' cannot be empty.")
	}

	 message := fmt.Sprintf("Hello, %v!", input)
	 return message, nil
}
