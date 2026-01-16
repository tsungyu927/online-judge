package main

fun main() {
  val solution = Solution()
  println(solution.removeElement(nums = intArrayOf(0, 1, 2, 2, 3, 0, 4, 2), `val` = 2))
  // 5 (Expect to be [0,1,4,0,3])
}

class Solution {
  fun removeElement(nums: IntArray, `val`: Int): Int {
    var k = 0
    var front = 0
    var end = nums.size - 1

    while (front <= end) {
      if (nums[front] == `val`) {
        nums[front] = nums[end]
        nums[end] = -1
        end--
      } else {
        front++
        k++
      }
    }
      
    return k
  }
}