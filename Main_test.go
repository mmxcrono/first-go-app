package main

import (
	"fmt"
	"testing"
	"strconv"
)

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestMain(t *testing.T) {

	fmt.Println("TestMain")
	testLanguages(t)

}

func testLanguages(t *testing.T) {
	languages := map[int]string{}
	languages[english] = "Hello"
	languages[spanish] = "Hola"
	languages[french] = "Bonjour"

	for key, value := range languages {
			
		t.Run("Language test for " + strconv.Itoa(key), func(t *testing.T) {
			got := Hello("John", key)
			want := value + ", John"

			assertCorrectMessage(t, got, want)

		})
	}


}