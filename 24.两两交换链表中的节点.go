/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	return swap(head)
}

func swap(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	nextHead := head.Next.Next
	next := head.Next
	next.Next = head
	head.Next = swap(nextHead)
	return next
}

// @lc code=end

