package solution

import (
	"unicode"
)

func isPalindrome(s string) bool {
	var newString []rune
	for _, c := range s {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			newString = append(newString, unicode.ToLower(c))
		}
	}

	for i, s := range newString {
		len := len(newString)
		if i > len/2 {
			return true
		}
		if s != newString[len-i-1] {
			return false
		}
	}
	return true
}
