/*
 * @lc app=leetcode.cn id=154 lang=golang
 *
 * [154] 寻找旋转排序数组中的最小值 II
 */

// @lc code=start
func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	start, end := 0, len(nums)-1
	for start+1 < end {
		for start < end && nums[end] == nums[end-1] {
			end--
		}
		for start < end && nums[start] == nums[start+1] {
			start++
		}
		midIndex := start + (end-start)/2
		if nums[midIndex] <= nums[end] {
			end = midIndex
		} else {
			start = midIndex
		}
	}
	if nums[start] <= nums[end] {
		return nums[start]
	} else {
		return nums[end]
	}
	return -1
}

// @lc code=end

