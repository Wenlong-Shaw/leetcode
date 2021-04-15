package leetcode

import (
	"math"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//递归的方法【带备忘录】（自上而下）
func CoinChange(coins []int, amount int) int {
	memo := make(map[int]int)

	var dp func(int) int
	dp = func(n int) int {
		if _, exists := memo[n]; exists {
			return memo[n]
		}
		if n == 0 {
			return 0
		}
		if n < 0 {
			return -1
		}
		res := math.MaxInt64
		for _, coin := range coins {
			subproblem := dp(n - coin)
			if subproblem == -1 {
				continue
			}
			res = min(res, 1+subproblem)
		}
		if res != math.MaxInt64 {
			memo[n] = res
			return memo[n]
		} else {
			return -1
		}
	}
	return dp(amount)
}

//迭代的方法（自底向上）
func CoinChange2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := range dp {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], 1+dp[i-coin])
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
