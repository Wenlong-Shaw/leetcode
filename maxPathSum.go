package leetcode

import "math"

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxPathSum(root *TreeNode) int {
	var ans int = math.MinInt64
	var maxSideSum func(*TreeNode) int
	maxSideSum = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := max(0, maxSideSum(root.Left))
		right := max(0, maxSideSum(root.Right))
		ans = max(ans, left+right+root.Val)
		return max(left, right) + root.Val
	}
	maxSideSum(root)
	return ans
}
