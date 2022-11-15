package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/Phogheus/GoLearnGo/greetings"
)

var (
	UnknownPathError ErrorResponse = ErrorResponse{"Unknown path."}
	InvalidMethodOnPathError ErrorResponse = ErrorResponse{"Invalid HTTP method on path."}
)

func main() {
	http.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = shiftPath(r.URL.Path)

	switch head {
	case "":
		serveHomePage(w)
	case "greeting":
		handleGreeting(w, r)
	case "favicon.ico":
		// TODO: Return image
	default:
		json.NewEncoder(w).Encode(UnknownPathError)
	}
}

func serveHomePage(w http.ResponseWriter) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func handleGreeting(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// pathTrail is the trail of the /greeting route; name will be the immediate next route variable
		name, _ := shiftPath(r.URL.Path) // Ignore any trailing route info
		greeting, err := greetings.Hello(name)

		if err != nil {
			errMsg := ErrorResponse{err.Error()}
			json.NewEncoder(w).Encode(errMsg)
		} else {
			greetingMsg := Greeting{name, greeting}
			json.NewEncoder(w).Encode(greetingMsg)
		}
	} else if r.Method == "POST" {
		var greeting NewGreeting
		err := json.NewDecoder(r.Body).Decode(&greeting)

		if err != nil {
			errMsg := ErrorResponse{err.Error()}
			json.NewEncoder(w).Encode(errMsg)
		} else {
			greetings.AddFormat(greeting.NewGreetingFormat)
			// return codes
		}
	} else {
		json.NewEncoder(w).Encode(InvalidMethodOnPathError)
	}
}

func shiftPath(p string) (head, tail string) {
    p = path.Clean("/" + p) // Prepend "/" to 'p' and clean - Clean will ensure proper formatting, allowing a "//" to become "/"
    i := strings.Index(p[1:], "/") + 1 // Find the next index of "/" after the prepended one (the trail)

	// If index is <= 0, there is no further trail
    if i <= 0 {
        return p[1:], "/"
    }

	// Slice path into head and trail (the rest)
    return p[1:i], p[i:]
}
