package leetcode

/*
* 220. 存在重复元素 III

* 给你一个整数数组 nums 和两个整数 k 和 t 。请你判断是否存在 两个不同下标 i 和 j，
* 使得 abs(nums[i] - nums[j]) <= t ，同时又满足 abs(i - j) <= k 。
* 如果存在则返回 true，不存在返回 false。

! 提示：
* 0 <= nums.length <= 2 * 10^4
* -2^31 <= nums[i] <= 2^31 - 1
* 0 <= k <= 10^4
* 0 <= t <= 2^31 - 1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/contains-duplicate-iii
*/

func abs(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	return false
}
