package main

import (
	"testing"
)

func TestIsSemiPalindrome(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Periodic repeated string (d=3)",
			input:    "abcabc",
			expected: true,
		},
		{
			name:     "Standard palindrome odd length (d=1)",
			input:    "racecar",
			expected: true,
		},
		{
			name:     "Standard palindrome even length (d=1)",
			input:    "abccba",
			expected: true,
		},
		{
			name:     "Both palindrome and periodic (d=1, d=3)",
			input:    "abaaba",
			expected: true,
		},
		{
			name:     "All identical characters",
			input:    "aaaa",
			expected: true,
		},
		{
			name:     "Alternating characters (d=2)",
			input:    "abababab",
			expected: true,
		},
		{
			name:     "Non-palindrome prime length",
			input:    "abcde",
			expected: false,
		},
		{
			name:     "Arbitrary non-palindrome even length",
			input:    "abcdef",
			expected: false,
		},
		{
			name:     "Two identical characters",
			input:    "aa",
			expected: true,
		},
		{
			name:     "Two different characters",
			input:    "ab",
			expected: false,
		},
		{
			name:     "Three character palindrome",
			input:    "aba",
			expected: true,
		},
		{
			name:     "Three character non-palindrome",
			input:    "abc",
			expected: false,
		},
		{
			name:     "Single character (no proper divisor)",
			input:    "a",
			expected: false,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isSemiPalindrome(tc.input)
			if result != tc.expected {
				t.Errorf("isSemiPalindrome(%q) = %v; expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func BenchmarkIsSemiPalindrome(b *testing.B) {
	benchmarks := []struct {
		name  string
		input string
	}{
		{"ShortMatch", "abcabc"},
		{"LongMatch", "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"},
		{"NoMatch", "abcdefghijklmnopqrstuvwxyz1234567890!@#$%^&*()_+~"},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				isSemiPalindrome(bm.input)
			}
		})
	}
}
