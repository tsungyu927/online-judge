package main

fun main() {
    val solution = Solution()
    println(solution.generateMatrix(n = 3).contentDeepToString())
    // [[1,2,3],[8,9,4],[7,6,5]]
}

enum class Direction {
    LEFT,
    RIGHT,
    TOP,
    BOTTOM,
}

class Solution {
    fun generateMatrix(n: Int): Array<IntArray> {
        val matrix: Array<IntArray> = Array(size = n) { _ -> IntArray(size = n) }
        var boundLeft = 0
        var boundRight = n - 1
        var boundTop = 1
        var boundBottom = n - 1

        var direction = Direction.RIGHT
        var x = 0
        var y = 0

        for (i in 0..(n * n) - 1) {
            matrix[y][x] = i + 1

            when (direction) {
                Direction.LEFT -> {
                    x--
                    if (x == boundLeft) {
                        boundLeft += 1
                        direction = Direction.TOP
                    }
                }

                Direction.RIGHT -> {
                    x++
                    if (x == boundRight) {
                        boundRight -= 1
                        direction = Direction.BOTTOM
                    }
                }

                Direction.TOP -> {
                    y--
                    if (y == boundTop) {
                        boundTop += 1
                        direction = Direction.RIGHT
                    }
                }

                Direction.BOTTOM -> {
                    y++
                    if (y == boundBottom) {
                        boundBottom -= 1
                        direction = Direction.LEFT
                    }
                }
            }
        }

        return matrix
    }
}
