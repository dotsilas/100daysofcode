package testing

import (
	"testing"
)

func TestGreet(t *testing.T) {
	name := "Bob"
	want := "Hello Bob"
	got, err := Greet(name)
	if got != want || err != nil {
		t.Fatal("TestGreet(%s): got %q/%v, want %q/nil", name, got, err, want)
	}
}
