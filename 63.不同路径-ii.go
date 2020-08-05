/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
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
		if obstacleGrid[i][0] == 1 || grid[i-1][0] == 0 {
			grid[i][0] = 0
		}
	}
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 1 || grid[0][j-1] == 0 {
			grid[0][j] = 0
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				grid[i][j] = 0
				continue
			}
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}
	return grid[m-1][n-1]
}

// @lc code=end

