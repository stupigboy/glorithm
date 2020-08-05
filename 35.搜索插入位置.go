/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start+1 < end {
		midIndex := start + (end-start)/2
		if nums[midIndex] == target {
			start = midIndex
		} else if nums[midIndex] > target {
			start = midIndex
		} else if nums[midIndex] < target {
			end = midIndex
		}
	}
	if nums[start] >= target {
		return start
	}
	if nums[end] >= target {
		return end
	}
	if nums[end] < target {
		return end + 1
	}
	return 0
}

// @lc code=end

