package main

import (
	"reflect"
	"testing"
)

// Helper function to convert slice to linked list
func sliceToList(nums []int) *ListNode {
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

// Helper function to convert linked list to slice
func listToSlice(head *ListNode) []int {
	var result []int
	current := head
	for current != nil {
			result = append(result, current.Val)
			current = current.Next
	}
	return result
}

func TestReverseList(t *testing.T) {
	testCases := []struct {
			input    []int
			expected []int
	}{
			{input: []int{1, 2, 3, 4, 5}, expected: []int{5, 4, 3, 2, 1}},
			{input: []int{1, 2}, expected: []int{2, 1}},
	}

	for _, tc := range testCases {
			t.Run("", func(t *testing.T) {
					inputList := sliceToList(tc.input)
					reversedList := reverseList(inputList)
					result := listToSlice(reversedList)
					if !reflect.DeepEqual(result, tc.expected) {
							t.Errorf("expected %v, got %v", tc.expected, result)
					}
			})
	}
}