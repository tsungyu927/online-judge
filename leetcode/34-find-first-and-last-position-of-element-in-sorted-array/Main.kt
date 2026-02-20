package main

fun main() {
    val nums = intArrayOf(5, 7, 7, 8, 8, 10)
    val target = 8
    val solution = Solution()
    println(solution.searchRange(nums, target).contentToString())
    // [3, 4]
}

class Solution {
    fun searchRange(
        nums: IntArray,
        target: Int,
    ): IntArray {
        var left = 0
        var right = nums.size - 1

        while (left <= right) {
            val mid = (left + right) / 2

            when {
                nums[mid] > target -> {
                    right = mid - 1
                }

                nums[mid] < target -> {
                    left = mid + 1
                }

                else -> {
                    var l = mid
                    var r = mid
                    var ansL = l
                    var ansR = r
                    while ((l != -1 && nums[l] == target) || (r < nums.size && nums[r] == target)) {
                        if (l != -1 && nums[l] == target) {
                            ansL = l
                            l--
                        }
                        if (r < nums.size && nums[r] == target) {
                            ansR = r
                            r++
                        }
                    }
                    return intArrayOf(ansL, ansR)
                }
            }
        }

        return intArrayOf(-1, -1)
    }
}
