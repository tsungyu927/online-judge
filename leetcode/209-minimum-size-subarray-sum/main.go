/*
Problem: Minimum Size Subarray Sum
Link: https://leetcode.com/problems/minimum-size-subarray-sum/

Solution: Use "Sliding Window"
*/
package main

import "math"

func minSubArrayLen(target int, nums []int) int {
	left := 0
	sum := 0
	result := math.MaxInt

	for right := 0; right < len(nums); right++ {
		sum += nums[right]

		for sum >= target {
			subLen := right - left + 1
			if subLen < result {
				result = subLen
			}
			sum -= nums[left]
			left++
		}
	}

	if result == math.MaxInt {
		return 0
	}
	return result
}
