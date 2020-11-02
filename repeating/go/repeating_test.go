package repeating

import (
	"testing"
)

func TestFirstNonRepeating(t *testing.T) {
	t.Run("An empty string returns ''", func(t *testing.T) {
		got := FirstNonRepeating("")
		want := ""
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("Returns 'a' when the input is 'a'", func(t *testing.T) {
		got := FirstNonRepeating("a")
		want := "a"
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("Returns 'a' when the input is 'ab'", func(t *testing.T) {
		got := FirstNonRepeating("ab")
		want := "a"
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("Returns 't' when input is 'stress'", func(t *testing.T) {
		got := FirstNonRepeating("stress")
		want := "t"
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("Returns '~' when input is 'aadede~'", func(t *testing.T) {
		got := FirstNonRepeating("aadede~")
		want := "~"
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("Ensures upper/lowercase are taken into account", func(t *testing.T) {
		got := FirstNonRepeating("Go hang a salami, I'm a lasagna hog!")
		want := ","
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}
