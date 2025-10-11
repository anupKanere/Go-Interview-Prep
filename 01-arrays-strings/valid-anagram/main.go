package main

import (
	"fmt"
	"sort"
)

/*
Problem: Valid Anagram
Difficulty: Easy
Source: LeetCode #242

Description:
Given two strings s and t, return true if t is an anagram of s, and false otherwise.
An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, 
typically using all the original letters exactly once.

Example 1:
Input: s = "anagram", t = "nagaram"
Output: true

Example 2:
Input: s = "rat", t = "car"
Output: false

Constraints:
- 1 <= s.length, t.length <= 5 * 10^4
- s and t consist of lowercase English letters only.

Time Complexity: O(n)
Space Complexity: O(1) - limited to 26 characters
*/

// Solution 1: Character Frequency Count - O(n) time, O(1) space
func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	
	// Count frequency of each character
	freq := make(map[rune]int)
	
	// Count characters in s
	for _, char := range s {
		freq[char]++
	}
	
	// Decrease count for characters in t
	for _, char := range t {
		freq[char]--
		if freq[char] < 0 {
			return false
		}
	}
	
	// Check if all counts are zero
	for _, count := range freq {
		if count != 0 {
			return false
		}
	}
	
	return true
}

// Solution 2: Sorting - O(n log n) time, O(1) space
func isAnagramSorting(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	
	// Convert to rune slices for sorting
	sRunes := []rune(s)
	tRunes := []rune(t)
	
	// Sort both slices
	sort.Slice(sRunes, func(i, j int) bool {
		return sRunes[i] < sRunes[j]
	})
	sort.Slice(tRunes, func(i, j int) bool {
		return tRunes[i] < tRunes[j]
	})
	
	// Compare sorted strings
	return string(sRunes) == string(tRunes)
}

// Solution 3: Array-based frequency count (optimized for lowercase letters)
func isAnagramArray(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	
	// Array for 26 lowercase letters
	freq := make([]int, 26)
	
	// Count characters in s
	for _, char := range s {
		freq[char-'a']++
	}
	
	// Decrease count for characters in t
	for _, char := range t {
		freq[char-'a']--
		if freq[char-'a'] < 0 {
			return false
		}
	}
	
	// Check if all counts are zero
	for _, count := range freq {
		if count != 0 {
			return false
		}
	}
	
	return true
}

// Solution 4: Using strings package (less efficient but readable)
func isAnagramStrings(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	
	// Sort both strings
	sSorted := sortString(s)
	tSorted := sortString(t)
	
	return sSorted == tSorted
}

// Helper function to sort a string
func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func main() {
	// Example usage
	testCases := []struct {
		s, t string
	}{
		{"anagram", "nagaram"},
		{"rat", "car"},
		{"hello", "hello"},
		{"aabbcc", "ccbbaa"},
	}
	
	for _, tc := range testCases {
		result := isAnagram(tc.s, tc.t)
		fmt.Printf("Is Anagram: s='%s', t='%s', result=%v\n", tc.s, tc.t, result)
	}
	
	fmt.Println("\nApproach comparison:")
	fmt.Println("1. Frequency Count: O(n) time, O(1) space - Best overall")
	fmt.Println("2. Array-based: O(n) time, O(1) space - Fastest for lowercase letters")
	fmt.Println("3. Sorting: O(n log n) time, O(1) space - Simple but slower")
	fmt.Println("4. Strings package: O(n log n) time, O(n) space - Most readable")
}