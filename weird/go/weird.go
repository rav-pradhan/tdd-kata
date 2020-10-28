package weird

import (
	"strings"
)

// ToWeirdCase converts a string into that dumb af uppercase-lowercase thing. E.g., "hello world" => "HeLlO WoRlD"
func ToWeirdCase(str string) string {
	substrings := strings.Split(str, " ")
	for index, word := range substrings {
		substrings[index] = weirdify(word)
	}
	return strings.Join(substrings, " ")
}

func weirdify(word string) string {
	weirdCaseWord := ""
	for index, char := range word {
		if isEven(index) {
			appendUppercaseLetterTo(&weirdCaseWord, char)
		} else {
			appendLowercaseLetterTo(&weirdCaseWord, char)
		}
	}
	return weirdCaseWord
}

func appendUppercaseLetterTo(str *string, letter rune) {
	*str = *str + strings.ToUpper(string(letter))
}

func appendLowercaseLetterTo(str *string, letter rune) {
	*str = *str + strings.ToLower(string(letter))
}

func isEven(pos int) bool {
	return pos%2 == 0
}
