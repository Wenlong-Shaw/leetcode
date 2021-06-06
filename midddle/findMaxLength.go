package leetcode

/*
! 和题目523.连续子数组的和是一摸一样的解法
* 525. 连续数组
* 给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。

* 示例 1:
* 输入: nums = [0,1]
* 输出: 2
* 说明: [0, 1] 是具有相同数量0和1的最长连续子数组。

* 示例 2:
* 输入: nums = [0,1,0]
* 输出: 2
* 说明: [0, 1] (或 [1, 0]) 是具有相同数量0和1的最长连续子数组。

* 提示：
* 1 <= nums.length <= 10……5
* nums[i] 不是 0 就是 1
* 通过次数31,617提交次数60,759
*/

func findMaxLength(nums []int) (ans int) {
	if len(nums) <= 1 {
		return 0
	}
	m := map[int]int{0: -1} //* 给那些前缀和是其倍数的相减用的
	presum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			presum -= 1
		} else {
			presum += nums[i]
		}
		if presumindex, ok := m[presum]; ok {
			if i-presumindex > ans {
				ans = i - presumindex
			}
		} else {
			m[presum] = i
		}
	}
	return ans
}
