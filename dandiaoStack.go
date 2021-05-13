package leetcode

func MonotoneStack(nums []int) []int {
	ans := make([]int, len(nums))
	var stack Stack
	for i := len(nums) - 1; i >= 0; i-- {
		for (!stack.IsEmpty()) && nums[stack.Top()] <= nums[i] { //小的数出栈，直到栈空，退出比较
			stack.Pop()
		}
		if stack.IsEmpty() { //如果比较到栈空才结束的话，说明他右边没有比他更大的，直接赋值-1
			ans[i] = -1
		} else {
			ans[i] = stack.Top() //找到了第一个比他大的数，记录下该较大数的index
		}
		stack.Push(i) //同时将比较数入栈
	}
	return ans
}
