/*
 * @lc app=leetcode.cn id=542 lang=golang
 *
 * [542] 01 矩阵
 */

// @lc code=start
func updateMatrix(matrix [][]int) [][]int {
	stack := make([][]int, 0)
	//j??0???
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				point := []int{i, j}
				stack = append(stack, point)
			} else {
				matrix[i][j] = -1
			}
		}
	}
	fourSide := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(stack) != 0 {
		point := stack[0]
		stack = stack[1:]
		for _, v := range fourSide {
			x := point[0] + v[0]
			y := point[1] + v[1]
			if x < 0 || y < 0 || x > len(matrix)-1 || y > len(matrix[0])-1 {
				continue
			}
			if matrix[x][y] == -1 {
				matrix[x][y] = matrix[point[0]][point[1]] + 1
				stack = append(stack, []int{x, y})
			}
		}

	}
	return matrix
}

// @lc code=end

