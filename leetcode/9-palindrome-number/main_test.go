package main

import "testing"

func TestPalindrome(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{121, true},   // Test case 1: Positive palindrome
		{-121, false}, // Test case 2: Negative number (not a palindrome)
		{123, false},  // Test case 3: Positive non-palindrome
	}

	for _, test := range tests {
		result := palindrome(test.input)
		if result != test.expected {
			t.Errorf(
				"palindrome(%d) = %v; want %v", test.input, result, test.expected,
			)
		}
	}
}
