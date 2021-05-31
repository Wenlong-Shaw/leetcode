package leetcode

import "math"

func isPowerOfFour(n int) bool {
	if n < 0 {
		return false
	}
	// TODO: max is 16, min is 0, mid = (max+min)/2
	min, max := 0, 16
	for min < max {
		mid := (min + max) / 2
		if n > int(math.Pow(float64(4), float64(mid))) {
			min = mid + 1
		} else {
			max = mid
		}
	}
	if n == int(math.Pow(float64(4), float64(min))) {
		return true
	}
	return false
}

func isPowerOfFour_1(n int) bool {
	if n == 1 {
		return true
	}
	if n < 0 || ((n & (n - 1)) == 0) {
		return false
	}
	return (n & 0x55555555) != 0
}
