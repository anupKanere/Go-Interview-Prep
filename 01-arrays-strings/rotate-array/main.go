package main

import "fmt"

func rotateArray(nums []int, k int) []int {
	n := len(nums)
	if n == 0 {
		return nums
	}

	// ensure k is within bounds
	k = k % n 

	// reverse entire array
	reverseArray(nums, 0, n-1)
	//reverse first k elements
	reverseArray(nums, 0, k-1)
	//reverse the remaining elements
	reverseArray(nums, k, n-1)

	return nums
}

func reverseArray(nums []int, left, right int) []int {

	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left, right = left+1, right-1
	}
	return nums
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	result := rotateArray(nums, k)
	fmt.Println(result)
}
