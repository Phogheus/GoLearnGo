package greetings

import "fmt"

func Hello(input string) string { // The function "Hello" is an "Exported Name" (accessible outside this package)
	 message := fmt.Sprintf("Hello, %v!", input)
	 return message
}
