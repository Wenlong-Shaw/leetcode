package leetcode

import "math"

/*
* 456. 132 模式
* 给你一个整数数组 nums ，数组中共有 n 个整数。132 模式的子序列 由三个整数 nums[i]、nums[j] 和 nums[k] 组成，
* 并同时满足：i < j < k 和 nums[i] < nums[k] < nums[j] 。
* 如果 nums 中存在 132 模式的子序列 ，返回 true ；否则，返回 false 。



进阶：很容易想到时间复杂度为 O(n^2) 的解决方案，你可以设计一个时间复杂度为 O(n logn) 或 O(n) 的解决方案吗？
*/

//时间复杂度为O(n)，枚举1，需要借助单调栈
func find132pattern(nums []int) bool {
	var s Stack
	n := len(nums)
	// arr := make([]int, n)
	k := math.MinInt64
	for i := n - 1; i >= 0; i-- {
		if nums[i] < k {
			return true
		}
		for !s.IsEmpty() && nums[s.Top()] < nums[i] { //! 注意是严格Pop出较小数
			//TODO: 记录Pop出的2的值，用于下次比较
			k = nums[s.Pop()]
		}
		if nums[i] > k {
			s.Push(i)
		}
	}
	return false
}

//* 官方解，使用数组切片，代替构建Stack数据结构。
//* 思路和步骤同上。
func find132pattern1(nums []int) bool {
	n := len(nums)
	candidateK := []int{nums[n-1]}
	maxK := math.MinInt64

	for i := n - 2; i >= 0; i-- {
		if nums[i] < maxK {
			return true
		}
		for len(candidateK) > 0 && nums[i] > candidateK[len(candidateK)-1] {
			maxK = candidateK[len(candidateK)-1]
			candidateK = candidateK[:len(candidateK)-1]
		}
		if nums[i] > maxK {
			candidateK = append(candidateK, nums[i])
		}
	}
	return false
}
