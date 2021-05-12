package leetcode

type MinStack struct {
	s   []int
	min []int
}

/** initialize your data structure here. */
func Constructor0() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.s = append(this.s, x)
	if len(this.min) == 0 {
		this.min = append(this.min, x)
		return
	}
	if this.min[len(this.min)-1] > x {
		this.min = append(this.min, x)
	} else {
		this.min = append(this.min, this.min[len(this.min)-1])
	}
}

func (this *MinStack) Pop() {
	this.s = this.s[:len(this.s)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	return this.s[len(this.s)-1]
}

func (this *MinStack) Min() int {
	return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
