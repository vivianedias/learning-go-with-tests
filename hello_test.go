package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Viviane")
	want := "Hello, Viviane"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}