package main

import (
	"reflect"
	"testing"
)

func TestRemoveElements(t *testing.T) {
	testCases := []struct {
		name     string
		head     *ListNode
		val      int
		expected *ListNode
	}{
		{
			name:     "Remove 6 from [1,2,6,3,4,5,6]",
			head:     createList([]int{1, 2, 6, 3, 4, 5, 6}),
			val:      6,
			expected: createList([]int{1, 2, 3, 4, 5}),
		},
		{
			name:     "Remove 1 from [1,1,1,1]",
			head:     createList([]int{1, 1, 1, 1}),
			val:      1,
			expected: nil,
		},
		{
			name:     "Remove 2 from [1,2,2,1]",
			head:     createList([]int{1, 2, 2, 1}),
			val:      2,
			expected: createList([]int{1, 1}),
		},
		{
			name:     "Remove 3 from empty list",
			head:     nil,
			val:      3,
			expected: nil,
		},
		{
			name:     "Remove 4 from [4,4,4,4]",
			head:     createList([]int{4, 4, 4, 4}),
			val:      4,
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := removeElements(tc.head, tc.val)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

// Helper function to create a linked list from a slice
func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, num := range nums[1:] {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return head
}