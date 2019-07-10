package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Ajan")
	want := "Hello, Ajan"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
