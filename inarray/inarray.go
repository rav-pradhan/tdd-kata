package inarray

import (
	"sort"
	"strings"
)

func InArray(substrings, words []string) []string {
	substringMatches := []string{}

	for _, substring := range substrings {
		checkForMatchingString(substring, words, &substringMatches)
	}
	sort.Strings(substringMatches)
	return substringMatches
}

func checkForMatchingString(substring string, words []string, matches *[]string) {
	for _, word := range words {
		if strings.Contains(word, substring) && !substringHasAlreadyBeenAdded(*matches, substring) {
			*matches = append(*matches, substring)
			break
		}
	}
}

func substringHasAlreadyBeenAdded(substringMatches []string, substring string) bool {
	for _, includedSubstring := range substringMatches {
		if substring == includedSubstring {
			return true
		}
	}
	return false
}
