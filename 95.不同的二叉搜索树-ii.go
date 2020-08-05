/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return generate(1, n)
}
func generate(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	var resp []*TreeNode
	for i := start; i <= end; i++ {
		left := generate(start, i-1)
		right := generate(i+1, end)
		for _, v := range left {
			for _, j := range right {
				node := &TreeNode{
					Val:   i,
					Left:  v,
					Right: j,
				}
				resp = append(resp, node)
			}
		}

	}
	return resp
}

// @lc code=end

