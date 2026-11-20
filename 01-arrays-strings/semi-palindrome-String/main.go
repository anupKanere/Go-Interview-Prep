package main

import "fmt"

// isSemiPalindrome checks whether a given string is a semi-palindrome.
//
// Definition:
// A string s of length n is called a semi-palindrome if there exists at least one
// proper divisor d of n (1 <= d < n, where n % d == 0) such that for every
// offset (0 <= offset < d), the subsequence formed by picking characters at
// step interval d:
//
//	s[offset], s[offset + d], s[offset + 2*d], ..., s[n - d + offset]
//
// forms a palindrome.
//
// Example:
//
//	For "abcabc" (length n = 6):
//	- Choose divisor d = 3 (proper divisor of 6):
//	  - offset 0: s[0], s[3] -> "aa" (palindrome)
//	  - offset 1: s[1], s[4] -> "bb" (palindrome)
//	  - offset 2: s[2], s[5] -> "cc" (palindrome)
//	Since all subsequences for d = 3 are palindromes, "abcabc" is a semi-palindrome.

func isSemiPalindrome(s string) bool {
	n := len(s)
	if n <= 1 {
		return false
	}

	// Find all proper divisors
	divisors := []int{}

	for d := 1; d <= n/d; d++ {
		if n%d == 0 {
			divisors = append(divisors, d)

			pair := n / d

			if pair != d && pair < n {
				divisors = append(divisors, pair)
			}
		}
	}

	// Check each divisor
	for _, d := range divisors {
		isValid := true

		// Check every group
		for offset := 0; offset < d; offset++ {
			left := offset
			right := n - d + offset

			// Two-pointer palindrome check
			for left < right {
				if s[left] != s[right] {
					isValid = false
					break
				}
				left += d
				right -= d
			}

			if !isValid {
				break
			}
		}

		if isValid {
			return true
		}

	}

	return false

}

func main() {
	str := "abcabc"
	result := "not a Semi-Palindrome String"
	if isSemiPalindrome(str) {
		result = "a Semi-Palindrome String"
	}
	fmt.Println("String", str, "is", result)
}
