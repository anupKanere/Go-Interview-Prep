package main

import (
	"reflect"
	"testing"
)

func TestRotateArray(t *testing.T) {
	testCases :=
		[]struct {
			name     string
			nums     []int
			k        int
			expected []int
		}{
			{
				name:     "Example 1",
				nums:     []int{1, 2, 3, 4, 5, 6, 7},
				k:        3,
				expected: []int{5, 6, 7, 1, 2, 3, 4},
			},
			{
				name:     "Example 2",
				nums:     []int{-1, -100, 3, 99},
				k:        2,
				expected: []int{3, 99, -1, -100},
			},
			{
				name:     "Example 3",
				nums:     []int{1, 2, 3, 4, 5, 6, 7},
				// first iteration -[7,1,2,3,4,5,6] 
				// second iteration -[6,7,1,2,3,4,5]
				// third iteration -[5,6,7,1,2,3,4]
				// fourth iteration -[4,5,6,7,1,2,3]
				// fifth iteration -[3,4,5,6,7,1,2]
				// sixth iteration -[2,3,4,5,6,7,1]
				// seventh iteration -[1,2,3,4,5,6,7]
				// eighth iteration -[7,1,2,3,4,5,6]
				// ninth iteration -[6,7,1,2,3,4,5]
				// tenth iteration -[5,6,7,1,2,3,4]
				// k = 10 % 7 = 3
				k:        10,
				expected: []int{5, 6, 7, 1, 2, 3, 4},
			},
		}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := rotateArray(tc.nums, tc.k)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
