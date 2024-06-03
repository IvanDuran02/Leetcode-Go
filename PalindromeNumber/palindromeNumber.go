package palindromenumber

import "fmt"

// Number must read the same from right to left and left to right

func PalindromeNumber(x int) bool {
	isPalindrome := true
	// Convert the integer to a string
	num := fmt.Sprintf("%d", x)
	// Convert the string to a slice of runes
	runes := []rune(num)
	runes_len := len(num) / 2

	for i := 0; i < runes_len; i++ {
		if runes[i] != runes[len(runes)-1-i] {
			isPalindrome = false
			break
		}
	}

	return isPalindrome
}
