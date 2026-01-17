package main

fun main() {
  val solution = Solution()
  println(solution.minSubArrayLen(target = 7, nums = intArrayOf(2, 3, 1, 2, 4, 3)))
  // 2
}

class Solution {
  fun minSubArrayLen(target: Int, nums: IntArray): Int {
    var left = 0
    var right = 0
    var acc = 0
    var minLen = Int.MAX_VALUE

    while (right < nums.size) {
      acc += nums[right]

      while (acc >= target) {
        val len = right - left + 1
        minLen = if (len < minLen) len else minLen
        acc -= nums[left]
        left++
      }
      right++
    }

    return if (minLen == Int.MAX_VALUE) 0 else minLen
  }
}