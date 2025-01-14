/*
Problem: Binary Search
Link: https://leetcode.com/problems/binary-search/
*/
package main

func binarySearch(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for start <= end {
		i := (start + end) / 2

		if nums[i] > target {
			end = i - 1
		} else if nums[i] < target {
			start = i + 1
		} else {
			return i
		}
	}

	return -1
}
