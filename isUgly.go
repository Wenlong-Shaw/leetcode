package leetcode

/*
* 给你一个整数 n ，请你判断 n 是否为 丑数 。如果是，返回 true ；否则，返回 false 。
* 丑数 就是只包含质因数 2、3 和/或 5 的正整数。


! 提示：
* -2^31 <= n <= 2^31 - 1
! 1 通常被视为丑数
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ugly-number
*/

func isUgly(n int) bool {
	var flag int
	if n < 1 {
		return false
	}
	for n > 0 {
		flag = n
		if n%2 == 0 {
			n = n / 2
		}
		if n%3 == 0 {
			n = n / 3
		}
		if n%5 == 0 {
			n = n / 5
		}
		if n == 1 {
			return true
		}
		if flag == n {
			return false
		}
	}
	return false
}
