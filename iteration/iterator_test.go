package iteration

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	got := Iterator("a", 5)
	want := "aaaaa"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func BenchmarkIterator(b *testing.B) {
	for b.Loop() {
		Iterator("a", 5)
	}
}

func ExampleIterator() {
	fmt.Println(Iterator("a", 5))
	// Output: "aaaaa"
}
