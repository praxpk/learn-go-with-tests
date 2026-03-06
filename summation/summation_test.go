package summation

import (
	"testing"
)

func TestSummation(t *testing.T) {
	got := Summation([]int{1, 2, 3, 4, 5})
	want := 15
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
