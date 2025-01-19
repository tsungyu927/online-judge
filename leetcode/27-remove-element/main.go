/*
Problem: Remove Element
Link: https://leetcode.com/problems/remove-element/

Solution: Use "Fast and Slow Pointers"
*/
package main

func removeElement(nums []int, val int) int {
	slow := 0

	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] == val {
			continue
		}
		nums[slow] = nums[fast]
		slow++
	}

	return slow
}
