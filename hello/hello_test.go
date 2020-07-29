package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("AJ")
	want := "Hello, AJ"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
