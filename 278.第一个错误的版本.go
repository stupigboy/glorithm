/*
 * @lc app=leetcode.cn id=278 lang=golang
 *
 * [278] 第一个错误的版本
 */

// @lc code=start
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	start, end := 0, n
	for start+1 < end {
		midIndex := start + (end-start)/2
		if isBadVersion(midIndex) {
			end = midIndex
		} else {
			start = midIndex
		}
	}
	if isBadVersion(start) {
		return start
	}
	if isBadVersion(end) {
		return end
	}
	return -1
}

// @lc code=end

