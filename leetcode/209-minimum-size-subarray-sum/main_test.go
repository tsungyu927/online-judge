package main

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "Leetcode testcase #1",
			nums:     []int{2, 3, 1, 2, 4, 3},
			target:   7,
			expected: 2,
		},
		{
			name:     "Leetcode testcase #2",
			nums:     []int{1, 4, 4},
			target:   4,
			expected: 1,
		},
		{
			name:     "Leetcode testcase #3",
			nums:     []int{1, 1, 1, 1, 1, 1, 1, 1},
			target:   11,
			expected: 0,
		},
		{
			name:     "First element is the answer",
			nums:     []int{10, 2, 3},
			target:   6,
			expected: 1,
		},
		{
			name:     "Answer is the length of nums",
			nums:     []int{1, 2, 3},
			target:   6,
			expected: 3,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(
			tc.name, func(t *testing.T) {
				t.Parallel()

				result := minSubArrayLen(tc.target, tc.nums)
				if result != tc.expected {
					t.Errorf(
						"minSubArrayLen(%d, %v) = %d; want %d",
						tc.target, tc.nums, result, tc.expected,
					)
				}
			},
		)
	}
}
