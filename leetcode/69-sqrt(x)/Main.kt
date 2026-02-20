package main

fun main() {
    val x = 8
    val solution = Solution()
    val result = solution.mySqrt(x)

    println(result)
    // 2
}

class Solution {
    fun mySqrt(x: Int): Int {
        var start = 0
        var end = x
        var ans = -1

        while (start <= end) {
            val mid = (start + end) / 2
            val midSqrt = mid.toLong() * mid.toLong()

            if (midSqrt > x.toLong()) {
                end = mid - 1
            } else if (midSqrt < x.toLong()) {
                ans = mid
                start = mid + 1
            } else {
                return mid
            }
        }

        return ans
    }
}
