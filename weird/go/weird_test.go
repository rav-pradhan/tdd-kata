package weird

import (
	"testing"
)

func TestWeird(t *testing.T) {
	t.Run("input of 'a' is uppercase converted", func(t *testing.T) {
		got := ToWeirdCase("a")
		want := "A"
		assertWeirdCase(t, got, want)
	})

	t.Run("input with an odd and even index character is properly converted", func(t *testing.T) {
		got := ToWeirdCase("ab")
		want := "Ab"
		assertWeirdCase(t, got, want)
	})

	t.Run("input of abc is converted to AbC", func(t *testing.T) {
		got := ToWeirdCase("abc")
		want := "AbC"
		assertWeirdCase(t, got, want)
	})

	t.Run("a string of ABC returns AbC", func(t *testing.T) {
		got := ToWeirdCase("ABC")
		want := "AbC"
		assertWeirdCase(t, got, want)
	})

	t.Run("a string of 'This is a test Looks like you passed' returns ThIs Is A TeSt LoOkS LiKe YoU PaSsEd", func(t *testing.T) {
		got := ToWeirdCase("This is a test Looks like you passed")
		want := "ThIs Is A TeSt LoOkS LiKe YoU PaSsEd"
		assertWeirdCase(t, got, want)
	})
}

func assertWeirdCase(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
