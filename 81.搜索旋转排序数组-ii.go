/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 */

// @lc code=start
func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
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
		if nums[midIndex] >= nums[start] {
			if nums[midIndex] >= target && nums[start] <= target {
				end = midIndex
			} else {
				start = midIndex
			}

		} else if nums[midIndex] < nums[end] {
			if nums[midIndex] <= target && nums[end] >= target {
				start = midIndex
			} else {
				end = midIndex
			}
		}
	}
	if nums[start] == target || nums[end] == target {
		return true
	}

	return false
}

// @lc code=end

