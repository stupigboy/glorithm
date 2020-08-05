/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	grid := make([][]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i] == nil {
				grid[i] = make([]int, n)
			}
			grid[i][j] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}
	return grid[m-1][n-1]
}

// @lc code=end

