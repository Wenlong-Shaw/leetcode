package leetcode

/*
* 87. 扰乱字符串

* 使用下面描述的算法可以扰乱字符串 s 得到字符串 t ：
* 	1.如果字符串的长度为 1 ，算法停止
* 	2.如果字符串的长度 > 1 ，执行下述步骤：
* 		1. 在一个随机下标处将字符串分割成两个非空的子字符串。即，如果已知字符串 s ，
*		则可以将其分成两个子字符串 x 和 y ，且满足 s = x + y 。
* 		2. 随机 决定是要「交换两个子字符串」还是要「保持这两个子字符串的顺序不变」。
*		即，在执行这一步骤之后，s 可能是 s = x + y 或者 s = y + x 。
* 		3. 在 x 和 y 这两个子字符串上继续从步骤 1 开始递归执行此算法。

* 给你两个 长度相等 的字符串 s1 和 s2，判断 s2 是否是 s1 的扰乱字符串。如果是，返回 true ；否则，返回 false 。

! 提示：
* s1.length == s2.length
* 1 <= s1.length <= 30
* s1 和 s2 由小写英文字母组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/scramble-string
*/
func isScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	n := len(s1)
	s1Map := make(map[byte][]int, len(s1))
	for i := range s1 {
		s1Map[s1[i]] = append(s1Map[s1[i]], i)
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i-1; j++ {
			// index, exists := s1Map[s2[j]]
			// if !exists {
			// 	return false
		}
	}
	return false
}
