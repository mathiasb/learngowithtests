package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to the world", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hellow to to people", func(t *testing.T) {
		got := Hello("You")
		want := "Hello, You!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf(
			"got %q want %q",
			got,
			want,
		)
	}
}
