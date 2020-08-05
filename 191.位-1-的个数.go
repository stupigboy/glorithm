/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] 位1的个数
 */

// @lc code=start
func hammingWeight(num uint32) int {
	var total int
	for num != 0 {
		num &= (num - 1)
		total++

	}
	return total
}

// @lc code=end

