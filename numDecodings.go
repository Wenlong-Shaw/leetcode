package leetcode

func numDecodings(s string) int {
	n := len(s)
	a, b, c := 0, 1, 0
	for i := 1; i <= n; i++ {
		c = 0
		if s[i] != '0' {
			c += b
		}
		if i > 1 && s[i] != '0' && ((s[i]*10 + s[i+1]) <= 26) {
			c += a
		}
		a, b = b, c
	}
	return c
}
