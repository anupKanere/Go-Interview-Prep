package main

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

/*
Problem: Group Anagrams
Difficulty: Medium
Source: LeetCode #49

Description:
Given an array of strings strs, group the anagrams together. You can return the answer in any order.
An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
typically using all the original letters exactly once.

Example 1:
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

Example 2:
Input: strs = [""]
Output: [[""]]

Example 3:
Input: strs = ["a"]
Output: [["a"]]

Constraints:
- 1 <= strs.length <= 10^4
- 0 <= strs[i].length <= 100
- strs[i] consists of lowercase English letters only.

Time Complexity: O(N * M log M) where N is number of strings, M is average length
Space Complexity: O(N * M)
*/

// Solution 1: Sort each string and use as key - O(N * M log M)
func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, str := range strs {
		// Sort the string to create a key
		sorted := sortString(str)
		groups[sorted] = append(groups[sorted], str)
	}

	// Convert map values to slice
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

// Solution 2: Use character frequency as key - O(N * M)
func groupAnagramsFrequency(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, str := range strs {
		// Create frequency key
		key := createFrequencyKey(str)
		groups[key] = append(groups[key], str)
	}

	// Convert map values to slice
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

// Solution 3: Using array for frequency (optimized for lowercase letters)
func groupAnagramsArray(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, str := range strs {
		// Create frequency key using array
		key := createFrequencyKeyArray(str)
		groups[key] = append(groups[key], str)
	}

	// Convert map values to slice
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

// Helper function to sort a string
func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

// Helper function to create frequency key
func createFrequencyKey(s string) string {
	freq := make(map[rune]int)
	for _, char := range s {
		freq[char]++
	}

	// Create key from frequency map
	var key strings.Builder
	for char := 'a'; char <= 'z'; char++ {
		if count, exists := freq[char]; exists {
			key.WriteRune(char)
			key.WriteString(fmt.Sprintf("%d", count))
		}
	}

	return key.String()
}

// Helper function to create frequency key using array
func createFrequencyKeyArray(s string) string {
	freq := make([]int, 26)
	for _, char := range s {
		freq[char-'a']++
	}

	// Create key from frequency array
	var key strings.Builder
	for i, count := range freq {
		if count > 0 {
			key.WriteByte(byte('a' + i))
			key.WriteString(fmt.Sprintf("%d", count))
		}
	}

	return key.String()
}

// Test cases
func TestGroupAnagrams(t *testing.T) {
	testCases := []struct {
		name     string
		strs     []string
		expected [][]string
	}{
		{
			name:     "Example 1",
			strs:     []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			name:     "Example 2 - Empty string",
			strs:     []string{""},
			expected: [][]string{{""}},
		},
		{
			name:     "Example 3 - Single character",
			strs:     []string{"a"},
			expected: [][]string{{"a"}},
		},
		{
			name:     "No anagrams",
			strs:     []string{"abc", "def", "ghi"},
			expected: [][]string{{"abc"}, {"def"}, {"ghi"}},
		},
		{
			name:     "All anagrams",
			strs:     []string{"abc", "bca", "cab"},
			expected: [][]string{{"abc", "bca", "cab"}},
		},
		{
			name:     "Mixed case",
			strs:     []string{"a", "b", "c", "a", "b"},
			expected: [][]string{{"a", "a"}, {"b", "b"}, {"c"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test sorting solution
			result := groupAnagrams(tc.strs)
			if !compareGroups(result, tc.expected) {
				t.Errorf("Sorting: Expected %v, got %v", tc.expected, result)
			}

			// Test frequency solution
			result2 := groupAnagramsFrequency(tc.strs)
			if !compareGroups(result2, tc.expected) {
				t.Errorf("Frequency: Expected %v, got %v", tc.expected, result2)
			}

			// Test array solution
			result3 := groupAnagramsArray(tc.strs)
			if !compareGroups(result3, tc.expected) {
				t.Errorf("Array: Expected %v, got %v", tc.expected, result3)
			}
		})
	}
}

// Helper function to compare two groups of anagrams
func compareGroups(groups1, groups2 [][]string) bool {
	if len(groups1) != len(groups2) {
		return false
	}

	// Sort each group and the overall groups
	for i := range groups1 {
		sort.Strings(groups1[i])
	}
	for i := range groups2 {
		sort.Strings(groups2[i])
	}

	// Sort the groups themselves
	sort.Slice(groups1, func(i, j int) bool {
		return groups1[i][0] < groups1[j][0]
	})
	sort.Slice(groups2, func(i, j int) bool {
		return groups2[i][0] < groups2[j][0]
	})

	// Compare
	for i := range groups1 {
		if len(groups1[i]) != len(groups2[i]) {
			return false
		}
		for j := range groups1[i] {
			if groups1[i][j] != groups2[i][j] {
				return false
			}
		}
	}

	return true
}

// Benchmark for performance testing
func BenchmarkGroupAnagrams(b *testing.B) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat", "abc", "bca", "cab"}
	for i := 0; i < b.N; i++ {
		groupAnagrams(strs)
	}
}

func BenchmarkGroupAnagramsArray(b *testing.B) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat", "abc", "bca", "cab"}
	for i := 0; i < b.N; i++ {
		groupAnagramsArray(strs)
	}
}

func main() {
	// Example usage
	testCases := [][]string{
		{"eat", "tea", "tan", "ate", "nat", "bat"},
		{"abc", "bca", "cab"},
		{"a", "b", "c"},
		{"", ""},
	}

	for _, strs := range testCases {
		result := groupAnagrams(strs)
		fmt.Printf("Group Anagrams: strs=%v, result=%v\n", strs, result)
	}

	fmt.Println("\nApproach comparison:")
	fmt.Println("1. Sorting: O(N * M log M) time - Simple but slower")
	fmt.Println("2. Frequency Map: O(N * M) time - Better for longer strings")
	fmt.Println("3. Frequency Array: O(N * M) time - Fastest for lowercase letters")
}
