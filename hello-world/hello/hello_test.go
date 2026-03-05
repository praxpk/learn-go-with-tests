package hello

import (
	"fmt"
	"testing"
)

func TestHell(t *testing.T) {
	name := "Chris"
	got := Hello(name)
	want := fmt.Sprintf("Hello, %s!", name)
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
