package main

import (
	"reflect"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Leetcode testcase #1",
			nums:     []int{-4, -1, 0, 3, 10},
			expected: []int{0, 1, 9, 16, 100},
		},
		{
			name:     "Leetcode testcase #2",
			nums:     []int{-7, -3, 2, 3, 11},
			expected: []int{4, 9, 9, 49, 121},
		},
		{
			name:     "All negative",
			nums:     []int{-10, -7, -4, -3, -1},
			expected: []int{1, 9, 16, 49, 100},
		},
		{
			name:     "All positive",
			nums:     []int{1, 3, 5, 7, 9},
			expected: []int{1, 9, 25, 49, 81},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(
			tc.name, func(t *testing.T) {
				t.Parallel()

				result := sortedSquares(tc.nums)
				if !reflect.DeepEqual(result, tc.expected) {
					t.Errorf(
						"sortedSquares(%v) = %v; want %v",
						tc.nums, result, tc.expected,
					)
				}
			},
		)
	}
}
