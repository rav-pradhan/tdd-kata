package longestrep

import (
	"fmt"
	"testing"
)

type Result struct {
	C rune
	L int
}

type Results []Result

func LongestRepetition(text string) Result {
	if text == "" {
		return Result{}
	}

	textAsSubstrings := splitTextByRepeatedCharacters(text)
	fmt.Println(textAsSubstrings)

	count := 1
	var char rune

	for index := range text {
		char = rune(text[index])
		if index > 0 && text[index] == text[index-1] {
			count++
			continue
		}
	}

	return Result{
		C: char,
		L: count,
	}
}

func splitTextByRepeatedCharacters(text string) []string {
	var substrings []string

	for index := range text {
		substring := text[index]
		if index > 0 && text[index] == text[index-1] {
			substring = substring + text[index]
			continue
		}

		substrings = append(substrings, substring)
	}
}

func TestLongestRepetition(t *testing.T) {
	t.Run("Empty string returns empty string", func(t *testing.T) {
		got := LongestRepetition("")
		want := Result{}
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("input of 'a' returns Result{C: 'a', L: 1}", func(t *testing.T) {
		got := LongestRepetition("a")
		want := Result{
			C: 'a',
			L: 1,
		}

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("input of 'aa' returns Result{'a', L: 2", func(t *testing.T) {
		got := LongestRepetition("aa")
		want := Result{
			C: 'a',
			L: 2,
		}

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("input of 'abbb' returns Result{'b', 3}", func(t *testing.T) {
		got := LongestRepetition("abbb")
		want := Result{
			C: 'b',
			L: 3,
		}

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("input of 'aaabcdefaabb' returns Result{'a', 3}", func(t *testing.T) {
		got := LongestRepetition("aaabcdefaabb")
		want := Result{
			C: 'a',
			L: 3,
		}

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}
