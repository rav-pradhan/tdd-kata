package leap

import "testing"

func TestIsLeapYear(t *testing.T) {
	t.Run("is not a leap year if not divisible by 4", func(t *testing.T) {
		want := false
		got := IsLeapYear(1997)
		assertEqualityBetween(t, got, want)
	})

	t.Run("is a leap year if divisible by 4", func(t *testing.T) {
		want := true
		got := IsLeapYear(1996)
		assertEqualityBetween(t, got, want)
	})

	t.Run("is a leap year if divisible by 400", func(t *testing.T) {
		want := true
		got := IsLeapYear(1600)
		assertEqualityBetween(t, got, want)
	})

	t.Run("is not a leap year if divisible by 100, but not by 400", func(t *testing.T) {
		want := false
		got := IsLeapYear(1800)
		assertEqualityBetween(t, got, want)
	})
}

func assertEqualityBetween(t *testing.T, got, want bool) {
	if got != want {
		t.Errorf("Got %t, wanted %t", got, want)
	}
}
