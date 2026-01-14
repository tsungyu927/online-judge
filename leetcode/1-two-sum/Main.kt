package two_sum

fun main() {
    val solution = Solution()
    println(solution.twoSum(intArrayOf(2, 7, 11, 15), 9).contentToString())
}

class Solution {
  fun twoSum(nums: IntArray, target: Int): IntArray {
    val diffMap = mutableMapOf<Int, Int>()

    nums.forEachIndexed { index, num -> 
      val expectNum = target - num

      diffMap[expectNum]?.let { return intArrayOf(it, index) }
      diffMap[num] = index
    }
    return intArrayOf()
  }
}