/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row := len(matrix)
	col := len(matrix[0])
	start, end := 0, row*col-1
	for start+1 < end {
		midIndex := start + (end-start)/2
		mid := matrix[midIndex/col][midIndex%col]
		if mid == target {
			return true
		} else if mid > target {
			end = midIndex
		} else if mid < target {
			start = midIndex
		}
	}
	if matrix[start/col][start%col] == target || matrix[end/col][end%col] == target {
		return true
	}
	return false

}

// @lc code=end

