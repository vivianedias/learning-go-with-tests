package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Viviane", "")
		want := "Hello, Viviane"
	
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
	
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Camila", "Spanish")
		want := "Hola, Camila"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Dani", "French")
		want := "Bonjour, Dani"
		assertCorrectMessage(t, got, want)
	})
}