package leetcode

/*
* 456. 132 模式
* 给你一个整数数组 nums ，数组中共有 n 个整数。132 模式的子序列 由三个整数 nums[i]、nums[j] 和 nums[k] 组成，
* 并同时满足：i < j < k 和 nums[i] < nums[k] < nums[j] 。
* 如果 nums 中存在 132 模式的子序列 ，返回 true ；否则，返回 false 。



进阶：很容易想到时间复杂度为 O(n^2) 的解决方案，你可以设计一个时间复杂度为 O(n logn) 或 O(n) 的解决方案吗？
*/

//时间复杂度为O(n)，枚举1，需要借助单调栈
func find132pattern(nums []int) bool {
	return false
}
