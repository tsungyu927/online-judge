package main

import kotlin.math.abs

fun main() {
  val solution = Solution()
  println(solution.sortedSquares(nums = intArrayOf(-4, -1, 0, 3, 10)).contentToString())
  // [0, 1, 9, 16, 100]
}

class Solution {
  fun sortedSquares(nums: IntArray): IntArray {
    var left = 0
    var right = nums.size - 1
    var index = nums.size - 1
    val result = IntArray(nums.size)

    while (left <= right) {
      if (abs(nums[left]) > abs(nums[right])) {
        result[index] = nums[left] * nums[left]
        left++
      } else {
        result[index] = nums[right] * nums[right]
        right--
      }
      
      index--
    }

    return result
  }
}