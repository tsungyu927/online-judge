/*
	Problem: Two Sum
	Link: https://leetcode.com/problems/two-sum/description/
*/

package main

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		expect := target - num

		if idx, found := numMap[expect]; found {
			return []int{idx, i}
		}

		// Store current num & i to numMap
		numMap[num] = i
	}

	return nil
}
