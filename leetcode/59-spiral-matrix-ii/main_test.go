package main

import (
	"reflect"
	"testing"
)

func TestGenerateMatrix(t *testing.T) {
	testCases := []struct{
		name string
		n int
		expected [][]int
	}{
		{
			name: "Leetcode testcase #1",
			n: 3,
			expected: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
		{
			name: "Leetcode testcase #2",
			n: 1,
			expected: [][]int{
				{1},
			},
		},
		{
			name: "Test n = 4",
			n: 4,
			expected: [][]int{
				{1, 2, 3, 4},
				{12, 13, 14, 5},
				{11, 16, 15, 6},
				{10, 9, 8, 7},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result := generateMatrix(tc.n)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}