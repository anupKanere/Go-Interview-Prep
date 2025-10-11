package main

import (
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "Example 1 - Has duplicates",
			nums:     []int{1, 2, 3, 1},
			expected: true,
		},
		{
			name:     "Example 2 - No duplicates",
			nums:     []int{1, 2, 3, 4},
			expected: false,
		},
		{
			name:     "Example 3 - Multiple duplicates",
			nums:     []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			expected: true,
		},
		{
			name:     "Single element",
			nums:     []int{1},
			expected: false,
		},
		{
			name:     "Two identical elements",
			nums:     []int{1, 1},
			expected: true,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: false,
		},
		{
			name:     "Negative numbers",
			nums:     []int{-1, -2, -1},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test hash set solution
			result := containsDuplicate(tc.nums)
			if result != tc.expected {
				t.Errorf("Hash Set: Expected %v, got %v", tc.expected, result)
			}
			
			// Test sorting solution
			result2 := containsDuplicateSorting(tc.nums)
			if result2 != tc.expected {
				t.Errorf("Sorting: Expected %v, got %v", tc.expected, result2)
			}
			
			// Test count solution
			result3 := containsDuplicateCount(tc.nums)
			if result3 != tc.expected {
				t.Errorf("Count: Expected %v, got %v", tc.expected, result3)
			}
		})
	}
}

// Benchmark for performance testing
func BenchmarkContainsDuplicate(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1}
	for i := 0; i < b.N; i++ {
		containsDuplicate(nums)
	}
}

func BenchmarkContainsDuplicateSorting(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1}
	for i := 0; i < b.N; i++ {
		containsDuplicateSorting(nums)
	}
}
