package leetcode

type CQueue struct {
	in []int
}

func Constructor_0() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.in = append(this.in, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.in) == 0 {
		return -1
	}
	a := this.in[0]
	this.in = this.in[1:]
	return a
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
