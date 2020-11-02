package repeating

import "strings"

func FirstNonRepeating(str string) string {
	charactersStore := charactersStore{
		characters: make(map[rune]int),
	}
	for _, char := range str {
		charactersStore.set(char)
	}
	return charactersStore.getFirstNonRepeatingCharacter()
}

type charactersStore struct {
	characters map[rune]int
	keys       []rune
}

func (c *charactersStore) set(char rune) {
	for _, val := range c.keys {
		if strings.ToLower(string(val)) == strings.ToLower(string(char)) {
			c.characters[val]++
			return
		}
	}
	c.keys = append(c.keys, char)
	c.characters[char] = 1
}

func (c *charactersStore) getFirstNonRepeatingCharacter() string {
	for _, val := range c.keys {
		if c.characters[val] == 1 {
			return string(val)
		}
	}
	return ""
}
