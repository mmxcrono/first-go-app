package main

import (
	"testing"
)

func TestMain(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Hello man", func(t *testing.T) {
		got := Hello("man")
		want := "Hello, man"

		assertCorrectMessage(t, got, want)

	})

	t.Run("Hello World when empty string passed", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)

	})


}
