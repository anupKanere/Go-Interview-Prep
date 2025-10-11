package main

import (
	"fmt"
	"sort"
)

/*
Problem: Contains Duplicate
Difficulty: Easy
Source: LeetCode #217

Description:
Given an integer array nums, return true if any value appears at least twice in the array, 
and return false if every element is distinct.

Example 1:
Input: nums = [1,2,3,1]
Output: true

Example 2:
Input: nums = [1,2,3,4]
Output: false

Example 3:
Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true

Constraints:
- 1 <= nums.length <= 10^5
- -10^9 <= nums[i] <= 10^9

Time Complexity: O(n)
Space Complexity: O(n)
*/

// Solution 1: Hash Set - O(n) time, O(n) space
func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	
	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	
	return false
}

// Solution 2: Sorting - O(n log n) time, O(1) space
func containsDuplicateSorting(nums []int) bool {
	sort.Ints(nums)
	
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	
	return false
}

// Solution 3: Using map with count - O(n) time, O(n) space
func containsDuplicateCount(nums []int) bool {
	count := make(map[int]int)
	
	for _, num := range nums {
		count[num]++
		if count[num] > 1 {
			return true
		}
	}
	
	return false
}

func main() {
	// Example usage
	testCases := [][]int{
		{1, 2, 3, 1},
		{1, 2, 3, 4},
		{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
	}
	
	for _, nums := range testCases {
		result := containsDuplicate(nums)
		fmt.Printf("Contains Duplicate: nums=%v, result=%v\n", nums, result)
	}
	
	fmt.Println("\nApproach comparison:")
	fmt.Println("1. Hash Set: O(n) time, O(n) space - Best for most cases")
	fmt.Println("2. Sorting: O(n log n) time, O(1) space - Best when space is limited")
	fmt.Println("3. Count Map: O(n) time, O(n) space - Similar to hash set")
}