package word2

import "unicode"

func unicodeLettersToLowerCase(s string) []rune {
	letters := make([]rune, 0)

	for _, rune_ := range s {
		if unicode.IsLetter(rune_) {
			letters = append(letters, unicode.ToLower(rune_))
		}
	}

	return letters
}

func IsPalindrome(s string) bool {
	letters := unicodeLettersToLowerCase(s)

	for i := range letters {
		if letters[i] != letters[len(letters)-i-1] {
			return false
		}
	}

	return true
}
