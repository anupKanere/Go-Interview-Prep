package main

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "Example 2",
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "Example 3",
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
		{
			name:     "Negative numbers",
			nums:     []int{-1, -2, -3, -4, -5},
			target:   -8,
			expected: []int{2, 4},
		},
		{
			name:     "Single solution",
			nums:     []int{1, 2, 3, 4, 5},
			target:   8,
			expected: []int{2, 4},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test hash map solution
			result := twoSum(tc.nums, tc.target)
			if len(result) != 2 {
				t.Errorf("Expected 2 indices, got %d", len(result))
			}
			if len(result) == 2 && tc.nums[result[0]]+tc.nums[result[1]] != tc.target {
				t.Errorf("Indices don't sum to target")
			}
		})
	}
}

func TestTwoSumBruteForce(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "Example 2",
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := twoSumBruteForce(tc.nums, tc.target)
			if len(result) != 2 {
				t.Errorf("Expected 2 indices, got %d", len(result))
			}
			if len(result) == 2 && tc.nums[result[0]]+tc.nums[result[1]] != tc.target {
				t.Errorf("Indices don't sum to target")
			}
		})
	}
}

// Benchmark for performance testing
func BenchmarkTwoSum(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 19
	for i := 0; i < b.N; i++ {
		twoSum(nums, target)
	}
}

func BenchmarkTwoSumBruteForce(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 19
	for i := 0; i < b.N; i++ {
		twoSumBruteForce(nums, target)
	}
}
