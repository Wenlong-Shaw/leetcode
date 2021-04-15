package leetcode

import (
	"sort"
)

/*
* 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？
* 请你找出所有和为 0 且不重复的三元组。
* 注意：答案中不可以包含重复的三元组。

! 提示：
! 0 <= nums.length <= 3000
! -10^5 <= nums[i] <= 10^5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
*/
func ThreeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := [][]int{}

	for i := 0; i < n; i++ {
		//第一重循环的前一个遍历过的元素，不能与下一个遍历的元素相同。【做去重处理】
		if i == 0 || nums[i] != nums[i-1] {
			k := n - 1
			target := -1 * nums[i]
			for j := i + 1; j < n; j++ {
				//
				if j > i+1 && nums[j] == nums[j-1] {
					continue
				}
				for j < k && nums[j]+nums[k] > target {
					k--
				}
				//
				if j == k {
					break
				}
				if nums[j]+nums[k] == target {
					ans = append(ans, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return ans
}
