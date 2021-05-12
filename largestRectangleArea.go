package leetcode

/*
* 84. 柱状图中最大的矩形
* 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
* 求在该柱状图中，能够勾勒出来的矩形的最大面积。

* 示例:
* 输入: [2,1,5,6,2,3]
* 输出: 10

 */
func largestRectangleArea(heights []int) int {
	var s, s1 Stack
	ans := 0
	n, lsmall, rsmall := len(heights), make([]int, len(heights)), make([]int, len(heights))
	for i := 0; i < n; i++ {
		//右边第一个小于数
		for !s.IsEmpty() && heights[s.Top()] >= heights[n-1-i] {
			s.Pop()
		}
		if s.IsEmpty() {
			rsmall[n-1-i] = n
		} else {
			rsmall[n-1-i] = s.Top()
		}
		s.Push(n - 1 - i)

		//左边第一个小于数
		for !s1.IsEmpty() && heights[s1.Top()] >= heights[i] {
			s1.Pop()
		}
		if s1.IsEmpty() {
			lsmall[i] = -1
		} else {
			lsmall[i] = s1.Top()
		}
		s1.Push(i)
	}
	for i := 0; i < n; i++ {
		ans = max(ans, heights[i]*(rsmall[i]-lsmall[i]-1))
	}
	return ans
}
