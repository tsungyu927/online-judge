package main

class NumArray(
    nums: IntArray,
) {
    val nums: IntArray

    init {
        this.nums = nums
    }

    fun sumRange(
        left: Int,
        right: Int,
    ): Int {
        var sum: Int = 0

        for (i in left..right) {
            sum += nums[i]
        }

        return sum
    }
}
