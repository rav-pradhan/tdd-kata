package consonant

import (
	"testing"
)

func TestSolve(t *testing.T) {
	t.Run("a string of consonants, 'xyz', returns value of 75", func(t *testing.T) {
		got := solve("xyz")
		want := 75
		assertSolve(t, got, want)
	})
	t.Run("vowels should not be included in the value calculation", func(t *testing.T) {
		got := solve("abc")
		want := 5
		assertSolve(t, got, want)
	})
	t.Run("random assortment of consonants and vowels should calculate only consonant values", func(t *testing.T) {
		got := solve("abababababfapeifapefijaefaepfjavnefjnfbhwyfnjsifjapnes")
		want := 143
		assertSolve(t, got, want)
	})
}

func assertSolve(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
