package main

fun main() {
  val solution = Solution()
  println(solution.search(nums = intArrayOf(-1, 0, 3, 5, 9, 12), target = 9))
  // 4
}

class Solution {
  fun search(nums: IntArray, target: Int): Int {
    var start = 0
    var end = nums.size - 1

    while (start <= end) {
      val mid = (start + end) / 2

      if (nums[mid] > target) {
        end = mid - 1
      } else if (nums[mid] < target) {
        start = mid + 1
      } else {
        return mid
      }
    }

    return -1
  }
}