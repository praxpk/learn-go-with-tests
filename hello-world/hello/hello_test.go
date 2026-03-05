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
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("Returning hello world if no name is give", func(t *testing.T) {
		want := "Hello, World!"
		got := Hello("")
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

}
