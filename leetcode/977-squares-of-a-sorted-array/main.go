/*
Problem: Squares of a Sorted Array
Link: https://leetcode.com/problems/squares-of-a-sorted-array/

Solution: Use "Two Pointers"
*/
package main

func sortedSquares(nums []int) []int {
	left := 0
	right := len(nums) - 1
	idx := len(nums) - 1

	result := make([]int, len(nums))

	for left <= right {
		if nums[left]*nums[left] > nums[right]*nums[right] {
			result[idx] = nums[left] * nums[left]
			left++
		} else {
			result[idx] = nums[right] * nums[right]
			right--
		}

		idx--
	}

	return result
}
