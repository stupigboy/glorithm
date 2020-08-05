/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */

// @lc code=start
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	start, end := 0, len(nums)-1
	for start <= end {
		midIndex := start + (end-start)/2
		if nums[midIndex] == target {
			return midIndex
		} else if nums[midIndex] > target {
			end = midIndex - 1
		} else if nums[midIndex] < target {
			start = midIndex + 1
		}
	}
	// if nums[start] == target {
	// 	return start
	// }
	// if nums[end] == target {
	// 	return end
	// }
	return -1
}

// @lc code=end

