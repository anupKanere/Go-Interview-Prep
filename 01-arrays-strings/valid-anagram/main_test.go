package main

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{
			name:     "Example 1 - Valid anagram",
			s:        "anagram",
			t:        "nagaram",
			expected: true,
		},
		{
			name:     "Example 2 - Not anagram",
			s:        "rat",
			t:        "car",
			expected: false,
		},
		{
			name:     "Same string",
			s:        "hello",
			t:        "hello",
			expected: true,
		},
		{
			name:     "Different lengths",
			s:        "a",
			t:        "ab",
			expected: false,
		},
		{
			name:     "Single character",
			s:        "a",
			t:        "a",
			expected: true,
		},
		{
			name:     "Empty strings",
			s:        "",
			t:        "",
			expected: true,
		},
		{
			name:     "Repeated characters",
			s:        "aabbcc",
			t:        "ccbbaa",
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test frequency count solution
			result := isAnagram(tc.s, tc.t)
			if result != tc.expected {
				t.Errorf("Frequency Count: Expected %v, got %v", tc.expected, result)
			}
			
			// Test sorting solution
			result2 := isAnagramSorting(tc.s, tc.t)
			if result2 != tc.expected {
				t.Errorf("Sorting: Expected %v, got %v", tc.expected, result2)
			}
			
			// Test array solution
			result3 := isAnagramArray(tc.s, tc.t)
			if result3 != tc.expected {
				t.Errorf("Array: Expected %v, got %v", tc.expected, result3)
			}
			
			// Test strings solution
			result4 := isAnagramStrings(tc.s, tc.t)
			if result4 != tc.expected {
				t.Errorf("Strings: Expected %v, got %v", tc.expected, result4)
			}
		})
	}
}

// Benchmark for performance testing
func BenchmarkIsAnagram(b *testing.B) {
	s, t := "anagram", "nagaram"
	for i := 0; i < b.N; i++ {
		isAnagram(s, t)
	}
}

func BenchmarkIsAnagramArray(b *testing.B) {
	s, t := "anagram", "nagaram"
	for i := 0; i < b.N; i++ {
		isAnagramArray(s, t)
	}
}

func BenchmarkIsAnagramSorting(b *testing.B) {
	s, t := "anagram", "nagaram"
	for i := 0; i < b.N; i++ {
		isAnagramSorting(s, t)
	}
}
