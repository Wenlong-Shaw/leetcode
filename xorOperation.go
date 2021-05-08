package leetcode

func xorOperation(n int, start int) int {
	rst := 0
	for i := 0; i < n; i++ {
		rst ^= (start + 2*i)
	}
	return rst
}
