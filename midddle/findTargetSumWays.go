package leetcode

/*
494. 目标和
给你一个整数数组 nums 和一个整数 target 。
向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。



示例 1：
输入：nums = [1,1,1,1,1], target = 3
输出：5
解释：一共有 5 种方法让最终目标和为 3 。
-1 + 1 + 1 + 1 + 1 = 3
+1 - 1 + 1 + 1 + 1 = 3
+1 + 1 - 1 + 1 + 1 = 3
+1 + 1 + 1 - 1 + 1 = 3
+1 + 1 + 1 + 1 - 1 = 3

示例 2：
输入：nums = [1], target = 1
输出：1


提示：
1 <= nums.length <= 20
0 <= nums[i] <= 1000
0 <= sum(nums[i]) <= 1000
-1000 <= target <= 100
*/

//* 回溯法，遍历每一种可能，将最后得到的sums进行target比较，查看是否一致。
func findTargetSumWays(nums []int, target int) int {
	ans := 0
	var sum func(int, int)
	sum = func(sums int, i int) {
		if i >= len(nums) {
			if sums == target {
				ans++
			}
			return
		}
		sum(sums+nums[i], i+1)
		sum(sums-nums[i], i+1)
	}
	sum(0, 0)
	return ans
}

//TODO: dp法。
//* 思路分析： 可将添加 “+”和“-”使表达式结果等于target，转化为正子集+负子集的和等于target
//* 于是就有了，sum(A)-sum(B)=target => 2sum(A) = target+sum(ALL) => sum(A) = (target+sum(ALL))/2
//* 从而转换为了01背包问题的解决思路。
func findTargetSumWays_dp(nums []int, target int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	//* 目标数过大和(sum-target)/2不能整除，则说明没有满足target的子集
	if sum < target || (sum-target)%2 == 1 {
		return 0
	}
	dp := make([][]int, len(nums)+1)
	for i := 0; i <= len(nums); i++ {
		dp[i] = make([]int, (sum-target)/2+1)
	}
	dp[0][0] = 1
	for i := 1; i <= len(nums); i++ {
		x := nums[i-1]
		for j := 0; j <= (sum-target)/2; j++ {
			dp[i][j] += dp[i-1][j]
			if j >= x {
				dp[i][j] += dp[i-1][j-x]
			}
		}
	}
	return dp[len(nums)][(sum-target)/2]
}
