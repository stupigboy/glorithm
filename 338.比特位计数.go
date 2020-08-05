/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] 比特位计数
 */

// @lc code=start
func countBits(num int) []int {
	result := make([]int, 0)
	for i := 0; i <= num; i++ {
		total := count(i)
		result = append(result, total)
	}
	return result
}
func count(n int) (total int) {
	for n != 0 {
		n &= (n - 1)
		total++
	}
	return total
}

// @lc code=end

