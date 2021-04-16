package leetcode

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

func robOnce(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, 2)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i%2] = max(nums[i]+dp[(i-2)%2], dp[(i-1)%2])
	}
	return dp[(n-1)%2]
}

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	return max(robOnce(nums[0:n-2]), robOnce(nums[1:n-1]))
}
