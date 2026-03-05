package hello

import "testing"

func TestHell(t *testing.T) {
	got := Hello()
	want := "Hello, World!"
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
