/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	start, end := 0, len(nums)-1
	for start+1 < end {
		midIndex := start + (end-start)/2
		if nums[midIndex] == target {
			return midIndex
		}
		if nums[start] <= nums[midIndex] {
			if nums[start] <= target && nums[midIndex] >= target {
				end = midIndex
			} else {
				start = midIndex
			}

		} else if nums[end] >= nums[midIndex] {
			if nums[midIndex] <= target && nums[end] >= target {
				start = midIndex
			} else {
				end = midIndex
			}
		}
	}

	if nums[start] == target {
		return start
	} else if nums[end] == target {
		return end
	}
	return -1
}

// @lc code=end

