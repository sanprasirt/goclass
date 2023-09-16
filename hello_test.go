package hello

import "testing"

func TestSayHello(t *testing.T) {
	want := "Hello, World"
	got := Say("World")
	if got != want {
		t.Errorf("Say() = %q, want %q", got, want)
	}
}