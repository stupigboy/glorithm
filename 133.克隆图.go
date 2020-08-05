/*
 * @lc app=leetcode.cn id=133 lang=golang
 *
 * [133] 克隆图
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	visited := make(map[*Node]*Node)
	return Clone(node, visited)
}
func Clone(node *Node, visited map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}
	if v, ok := visited[node]; ok {
		return v
	}
	newNode := &Node{Val: node.Val, Neighbors: make([]*Node, len(node.Neighbors))}
	visited[node] = newNode
	for i := 0; i < len(node.Neighbors); i++ {
		newNode.Neighbors[i] = Clone(node.Neighbors[i], visited)
	}
	return newNode
}

// @lc code=end
