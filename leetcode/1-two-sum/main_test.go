package main

import "testing"

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	expected := []int{0, 1}

	result := twoSum(nums, target)
	if len(result) != 2 || result[0] != expected[0] || result[1] != expected[1] {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
