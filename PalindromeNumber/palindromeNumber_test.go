package palindromenumber

import (
	"testing"
)

// TestIsPalindrome tests the isPalindrome function with multiple test cases
func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{121, true},
		{-121, false},
		{10, false},
		{12321, true},
		{0, true},
		{1, true},
		{1221, true},
		{123, false},
	}

	for _, test := range tests {
		result := PalindromeNumber(test.input)
		if result != test.expected {
			t.Errorf("isPalindrome(%d) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
