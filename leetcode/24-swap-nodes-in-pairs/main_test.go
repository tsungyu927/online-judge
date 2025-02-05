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

func TestSwapPairs(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{input: []int{1, 2, 3, 4}, output: []int{2, 1, 4, 3}},
		{input: []int{1, 2, 3}, output: []int{2, 1, 3}},
		{input: []int{1}, output: []int{1}},
		{input: []int{}, output: []int{}},
	}

	for _, test := range tests {
		inputList := sliceToList(test.input)
		expectedList := sliceToList(test.output)
		resultList := swapPairs(inputList)
		if !reflect.DeepEqual(listToSlice(resultList), listToSlice(expectedList)) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.output, listToSlice(resultList))
		}
	}
}