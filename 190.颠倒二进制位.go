/*
 * @lc app=leetcode.cn id=190 lang=golang
 *
 * [190] 颠倒二进制位
 */

// @lc code=start
func reverseBits(num uint32) uint32 {
	var result uint32
	pow := 31
	for num != 0 {
		//获取最后一位
		result += (num & 1) << pow
		num = num >> 1
		pow--
	}
	return result
}

// @lc code=end

