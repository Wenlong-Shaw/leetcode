package leetcode

//常规遍历法
func addDigits(num int) int {
	ans := 0
	for num > 10 {
		num1 := num % 10
		num = num / 10
		ans += num1
		if num < 10 {
			ans += num
			num = ans
			ans = 0
		}
	}
	if num < 10 {
		return num
	}
	return num
}

//时间复杂度为1
func addDigits_1(num int) int {
	return (num-1)%9 + 1
}
