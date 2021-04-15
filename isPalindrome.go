package leetcode

/*
* 9. 回文数
*	给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
* 	回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。

* 提示：
* 	-2^31 <= x <= 2^31 - 1

!进阶：你能不将整数转为字符串来解决这个问题吗？
*/
import (
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	str := strconv.Itoa(x)
	length := len(str)
	for i := 0; i <= length/2; i++ {
		if str[i] != str[length-i-1] {
			return false
		}
	}
	return false
}
