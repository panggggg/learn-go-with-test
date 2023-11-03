package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to Bangchan", func(t *testing.T) {
		got := Hello("Bangchan", "Korea")
		want := "Annyong, Bangchan"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Korea", func(t *testing.T) {
		got := Hello("Bangchan", "Korea")
		want := "Annyong, Bangchan"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Hyunjin", "French")
		want := "Bonjour, Hyunjin"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // tell the test suite assertCorrectMessage() is helper By doing this when it fails the line number reported will be in our function call rather than inside our test helper.
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
