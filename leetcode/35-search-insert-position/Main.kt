package main

fun main() {
    val solution = Solution()
    println(solution.searchInsert(nums = intArrayOf(1, 3, 5, 6), target = 5))
    // 2
}

class Solution {
    fun searchInsert(
        nums: IntArray,
        target: Int,
    ): Int {
        var left = 0
        var right = nums.size - 1

        while (left <= right) {
            val mid = (left + right) / 2

            when {
                nums[mid] > target -> right = mid - 1
                nums[mid] < target -> left = mid + 1
                else -> return mid
            }
        }

        return left
    }
}
