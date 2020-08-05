/*
 * @lc app=leetcode.cn id=344 lang=golang
 *
 * [344] 反转字符串
 */

// @lc code=start
func reverseString(s []byte) {
	ret := make([]byte, 0)
	reverse(0, s, &ret)
	for k, v := range ret {
		s[k] = v
	}
}

func reverse(i int, s []byte, ret *[]byte) {
	if i > len(s)-1 {
		return
	}
	reverse(i+1, s, ret)
	*ret = append(*ret, s[i])
	return
}

// @lc code=end

