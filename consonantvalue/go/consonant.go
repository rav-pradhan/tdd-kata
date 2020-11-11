package consonant

import (
	"strings"
)

var alphabetValues = map[rune]int{
	'a': 1, 'b': 2, 'c': 3, 'd': 4,
	'e': 5, 'f': 6, 'g': 7, 'h': 8,
	'i': 9, 'j': 10, 'k': 11, 'l': 12,
	'm': 13, 'n': 14, 'o': 15, 'p': 16,
	'q': 17, 'r': 18, 's': 19, 't': 20,
	'u': 21, 'v': 22, 'w': 23, 'x': 24,
	'y': 25, 'z': 26,
}

func solve(str string) int {
	largestValue := 0

	consonantSubstrings := strings.FieldsFunc(str, splitByVowels)
	for _, substring := range consonantSubstrings {
		substringValue := calculateSubstringValue(substring)
		if substringValue > largestValue {
			largestValue = substringValue
		}
	}
	return largestValue
}

func splitByVowels(char rune) bool {
	return isVowel(char)
}

func isVowel(char rune) bool {
	return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u'
}

func calculateSubstringValue(substring string) int {
	value := 0
	for _, char := range substring {
		value += alphabetValues[char]
	}
	return value
}
