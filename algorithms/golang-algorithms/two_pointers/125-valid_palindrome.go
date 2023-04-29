package main

import (
	"unicode"
)

func isAlphaNumeric(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}

// TC: O(n)
// SC: O(1) (Technically in golang it would be O(n) because converting striing to runes)
func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	str := []rune(s)

	for i <= j {
		left := unicode.ToLower(str[i])
		right := unicode.ToLower(str[j])

		// Alphanumeric checks
		if !isAlphaNumeric(left) {
			i++
			continue
		}

		if !isAlphaNumeric(right) {
			j--
			continue
		}

		// Check
		if left != right {
			return false
		}

		i++
		j--
	}

	return true
}

func main() {
	s := "A man, a plan, a canal: Panama"
	println(isPalindrome(s))
}
