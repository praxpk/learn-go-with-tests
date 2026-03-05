package hello

import (
	"fmt"
	"testing"
)

func TestHelloGreeting(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		name := "Chris"
		got := Hello(name)
		want := fmt.Sprintf("Hello, %s!", name)
		assertCorrectMessage(t, got, want)
	})
	t.Run("Returning hello world if no name is give", func(t *testing.T) {
		want := "Hello, World!"
		got := Hello("")
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
