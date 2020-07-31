package main

import (
	"bytes"
	"testing"
)

// Test fmt.Printf - prints to stdout,
// pretty hard to capture using the testing framework.
// What we need to do is to able to inject
// (which is just a fancy word for pass in)
// the dependency of printing.

// Our function doesn't need to care where or how the
// printing happens, so we should accepts an interface
// rather than a concrete type.

// If we do that we can change the implementation to print to
// something we control so that we can test it.
// In "real life" we would inject in something that writes to stdout.

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
