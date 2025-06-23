package hello_test

import "testing"

func TestHello(t *testing.T) {
	if got, want := "Hello, world!", Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func Hello() string {
	return "Hello, world!"
}
