package greetings

import (
	"regexp"
	"testing"
)

// Tests the "Hello" function under happy-path conditions
func TestHello(t *testing.T) {
	name := "My Awesome Name"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err) // Logs the error and then fails the test
	} else if !want.MatchString(msg) { // Learning note: "else if" won't compile if it's on a new line after the "if" closing bracket; same for "else"
		t.Fatal("Result was not formatted as expected.")
	}
}

// Tests the "Hello" function when sending an empty string
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")

	if err == nil {
		t.Fatal("Not receiving an error is unexpected.")
	} else if msg != "" {
		t.Fatal("Result contained a value, and is unexpected behavior.")
	}
}

// Tests the "MultipleHellos" function under happy-path conditions
func TestMultipleHellos(t *testing.T) {
	testNames := []string { "Awesome Name 1", "Awesome Name 2", "Awesome Name 3" }

	msgs, err := MultipleHellos(testNames)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	
	msgsLength := len(msgs)

	if msgsLength != len(testNames) {
		t.Fatalf("Unexpected length of mapped results: %d", msgsLength)
	}
	
	for i, testName := range testNames {
		want := regexp.MustCompile(`\b` + testName + `\b`)
		
		if !want.MatchString(testName) {
			t.Fatalf("Result at index %d was not formatted as expected (result = %v).", i, testName)
		}
	}
}

func TestMultipleHellosWithEmpty(t *testing.T) {
	testNames := []string { "Awesome Name 1", "", "Awesome Name 3" }

	_, err := MultipleHellos(testNames) // The '_' discards the returned "messages" map

	if err == nil {
		t.Fatal("Not receiving an error is unexpected.")
	}
}

func TestMultipleHellosWithDuplicates(t *testing.T) {
	testNames := []string { "Awesome Name 1", "Awesome Name 2", "Awesome Name 1" }

	msgs, err := MultipleHellos(testNames)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	
	msgsLength := len(msgs)

	if msgsLength != len(testNames) {
		t.Fatalf("Unexpected length of mapped results: %d", msgsLength)
	}
	
	for i, testName := range testNames {
		want := regexp.MustCompile(`\b` + testName + `\b`)
		
		if !want.MatchString(testName) {
			t.Fatalf("Result at index %d was not formatted as expected (result = %v).", i, testName)
		}
	}
}

func TestAddFormat(t *testing.T) {
	initialLengthOfFormats := len(formats)

	AddFormat("New format")

	newLengthOfFormats := len(formats)

	if newLengthOfFormats != initialLengthOfFormats + 1 {
		t.Fatal("Add new format failed.")
	}
}

func TestAddFormatEmpty(t *testing.T) {
	err := AddFormat("")

	if err == nil {
		t.Fatal("Not receiving an error is unexpected.")
	}
}
