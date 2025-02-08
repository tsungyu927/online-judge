package main

import (
	"reflect"
	"testing"
)

func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func sliceToList(nums []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return dummy.Next
}

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		head     []int
		n        int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{[]int{1, 2}, 1, []int{1}},
		{[]int{1, 2}, 2, []int{2}},
	}

	for _, test := range tests {
		head := sliceToList(test.head)
		result := removeNthFromEnd(head, test.n)
		resultSlice := listToSlice(result)
		if !reflect.DeepEqual(resultSlice, test.expected) {
			t.Errorf("For input %v and n = %d, expected %v but got %v", test.head, test.n, test.expected, resultSlice)
		}
	}
}