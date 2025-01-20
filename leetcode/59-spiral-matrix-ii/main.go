/*
	Problem: Spiral Matrix II
	Link: https://leetcode.com/problems/spiral-matrix-ii/
*/

package main

func generateMatrix(n int) [][]int {
  startRow, startCol := 0, 0
  num := 1
  
  round := n / 2
  loop := n - 1

  result := make([][]int, n)
  for i := 0; i < n; i++ {
    result[i] = make([]int, n)
  }

  for round > 0 {
    row, col := startRow, startCol

    for i := 0; i < loop; i++ {
      // Left -> Right
      result[row][col] = num
      col++
      num++
    }
    for i := 0; i < loop; i++ {
      // Top -> Bottom
      result[row][col] = num
      row++
      num++
    }
    for i := 0; i < loop; i++ {
      // Right -> Left
      result[row][col] = num
      col--
      num++
    }
    for i := 0; i < loop; i++ {
      // Bottom -> Top
      result[row][col] = num
      row--
      num++
    }

    startRow++
    startCol++
    loop-=2
    round--
  }

  if n % 2 == 1 {
    // Set the center value
    result[n / 2][n / 2] = num
  }

  return result
}
