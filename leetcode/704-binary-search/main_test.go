package main

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected int
	}{
		{
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   9,
			expected: 4,
		},
		{
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   2,
			expected: -1,
		},
	}

	for _, test := range tests {
		t.Run(
			"Test Binary Search", func(t *testing.T) {
				got := binarySearch(test.nums, test.target)
				if got != test.expected {
					t.Errorf(
						"binarySearch(%v, %d) = %d; want %d",
						test.nums, test.target, got, test.expected,
					)
				}
			},
		)

	}
}
