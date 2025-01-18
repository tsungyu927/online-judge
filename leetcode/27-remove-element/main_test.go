package main

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	testCases := []struct {
		name          string
		nums          []int
		val           int
		expectedLen   int
		expectedSlice []int
	}{
		{
			name:          "No occurrences of val",
			nums:          []int{1, 2, 3, 4, 5},
			val:           6,
			expectedLen:   5,
			expectedSlice: []int{1, 2, 3, 4, 5},
		},
		{
			name:          "All elements are val",
			nums:          []int{2, 2, 2},
			val:           2,
			expectedLen:   0,
			expectedSlice: []int{}, // first 0 elements
		},
		{
			name:          "Some occurrences",
			nums:          []int{3, 2, 2, 3},
			val:           3,
			expectedLen:   2,
			expectedSlice: []int{2, 2}, // first 2 elements after removal
		},
		{
			name:          "Empty array",
			nums:          []int{},
			val:           1,
			expectedLen:   0,
			expectedSlice: []int{},
		},
	}

	for _, tc := range testCases {
		// Capture range variable
		tc := tc
		t.Run(
			tc.name, func(t *testing.T) {
				t.Parallel() // Optional: run tests in parallel

				// Make a copy if we need the original unmodified, but here we can modify in place
				numsCopy := make([]int, len(tc.nums))
				copy(numsCopy, tc.nums)

				gotLen := removeElement(numsCopy, tc.val)

				if gotLen != tc.expectedLen {
					t.Errorf(
						"For input %v, val=%d, got length = %d, want %d",
						tc.nums, tc.val, gotLen, tc.expectedLen,
					)
				}

				// Check the first gotLen elements of numsCopy
				if !reflect.DeepEqual(numsCopy[:gotLen], tc.expectedSlice) {
					t.Errorf(
						"After removal, got slice = %v, want %v",
						numsCopy[:gotLen], tc.expectedSlice,
					)
				}
			},
		)
	}
}
