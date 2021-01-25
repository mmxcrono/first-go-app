package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("Hello man", func(t *testing.T) {
		got := Hello("man")
		want := "Hello, man"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
