package leetcode

func maxSubArray(nums []int) int {
	a, b, c, n := nums[0], 0, nums[0], len(nums)
	for r := 1; r < n; r++ {
		b = nums[r]
		b = max(b, a+b)
		c = max(c, b)
		a = b
	}
	return c
}
