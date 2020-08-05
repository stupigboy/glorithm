/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lc code=start
func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	start, end := 0, len(nums)-1
	for start+1 < end {
		midIndex := start + (end-start)/2
		if nums[midIndex] >= nums[end] {
			start = midIndex
		} else {
			end = midIndex
		}
	}
	if nums[start] <= nums[end] {
		return nums[start]
	} else {
		return nums[end]
	}
}

// @lc code=end

