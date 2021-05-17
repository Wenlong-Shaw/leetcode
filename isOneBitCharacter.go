package leetcode

func isOneBitCharacter(bits []int) bool {
	n := len(bits)
	if n == 1 {
		return true
	}
	for i := 0; i < n; i++ {
		if bits[i] == 1 {
			i += 1
			if i == n-1 {
				return false
			}
		}
		if bits[i] == 0 {
			if i == n-1 {
				return true
			}
		}
	}
	return false
}
