/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

// @lc code=start
func fib(N int) int {
	return dfs(N)
}

var m map[int]int = make(map[int]int)

func dfs(n int) int {
	if n < 2 {
		return n
	}
	if m[n] != 0 {
		return m[n]
	}
	res := dfs(n-1) + dfs(n-2)
	m[n] = res
	return res

}

// @lc code=end

