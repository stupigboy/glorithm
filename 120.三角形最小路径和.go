package glorithm

/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {
	l := len(triangle)
	f := make([][]int, l)
	for i := 0; i <= l-1; i++ {
		if f[i] == nil {
			f[i] = make([]int, len(triangle[i]))
		}
		for j := 0; j < len(triangle[i]); j++ {
			f[i][j] = triangle[i][j]
		}
	}
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			f[i][j] = min(f[i+1][j], f[i+1][j+1]) + triangle[i][j]
		}
	}
	return f[0][0]
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end
