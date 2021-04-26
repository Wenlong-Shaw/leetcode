package leetcode

/*
* 300. 最长递增子序列
* 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
* 子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。
* 例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

* 例 1：
* 输入：nums = [10,9,2,5,3,7,101,18]
* 输出：4
* 解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。

* 示例 2：
* 输入：nums = [0,1,0,3,2,3]
* 输出：4

* 示例 3：
* 输入：nums = [7,7,7,7,7,7,7]
* 输出：1


! 提示：
* 1 <= nums.length <= 2500
* -10^4 <= nums[i] <= 10^4

! 进阶：
* 你可以设计时间复杂度为 O(n2) 的解决方案吗？
* 你能将算法的时间复杂度降低到 O(n log(n)) 吗?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-increasing-subsequence
*/

//动态规划求解：每前进一个状态，都需要遍历前面的所有状态，看看是否有符合的状态。
func lengthOfLIS(nums []int) int {
	result := 0
	n := len(nums)
	if n == 0 {
		return result
	}
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i < n; i++ {
		//初始化dp[i]应该为1，避免直接继承前者，导致存在上下坡度时，不是严格增序。
		dp[i] = 1
		//每次更新dp[i], 都需要从头遍历每一个dp[j]状态是否符合递增，最后选出最大得数
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		//result 保存者全局的最大值
		result = max(dp[i], result)
	}
	return result
}
