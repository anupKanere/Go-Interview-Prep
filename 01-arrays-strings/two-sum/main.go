package main

import (
	"fmt"
)

/*
Problem: Two Sum
Difficulty: Easy
Source: LeetCode #1

Description:
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

Example 1:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:
Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:
Input: nums = [3,3], target = 6
Output: [0,1]

Constraints:
- 2 <= nums.length <= 10^4
- -10^9 <= nums[i] <= 10^9
- -10^9 <= target <= 10^9
- Only one valid answer exists.

Time Complexity: O(n)
Space Complexity: O(n)
*/

// Solution 1: Brute Force - O(n²)
func twoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// Solution 2: Hash Map - O(n)
func twoSum(nums []int, target int) []int {
	// Map to store number and its index
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if index, exists := numMap[complement]; exists {
			return []int{index, i}
		}
		numMap[num] = i
	}

	return []int{}
}

func main() {
	// Example usage
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Printf("Two Sum: nums=%v, target=%d, result=%v\n", nums, target, result)

	// Performance comparison
	fmt.Println("\nPerformance comparison:")
	fmt.Println("Hash Map approach is O(n) time complexity")
	fmt.Println("Brute Force approach is O(n²) time complexity")
}
