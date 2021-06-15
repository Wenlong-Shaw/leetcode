package leetcode

/* *[01背包问题]
* 1049. 最后一块石头的重量 II
* 有一堆石头，用整数数组 stones 表示。其中 stones[i] 表示第 i 块石头的重量。
* 每一回合，从中选出任意两块石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下：
* 如果 x == y，那么两块石头都会被完全粉碎；
* 如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。
* 最后，最多只会剩下一块 石头。返回此石头 最小的可能重量 。如果没有石头剩下，就返回 0。


* 示例 1：
* 输入：stones = [2,7,4,1,8,1]
* 输出：1
* 解释：
* 组合 2 和 4，得到 2，所以数组转化为 [2,7,1,8,1]，
* 组合 7 和 8，得到 1，所以数组转化为 [2,1,1,1]，
* 组合 2 和 1，得到 1，所以数组转化为 [1,1,1]，
* 组合 1 和 1，得到 0，所以数组转化为 [1]，这就是最优值。

* 示例 2：
* 输入：stones = [31,26,33,21,40]
* 输出：5

* 示例 3：
* 输入：stones = [1,2]
* 输出：1

! 提示：
* 1 <= stones.length <= 30
* 1 <= stones[i] <= 100
*/
//TODO: 问题可以转化为在stones的各个元素前添加 “+” 或 “-” 最后使得表达式的值最小为多少？
//* 则转化为了最基本的01背包问题，将stones的元素和分为 SumA和SumB两部分，则可知当 SumA -SumB = 0时表达式结果最小，
//* 即 SumA == SumB == Sum/2 时，结果最小。即可以转化为，是否选择将 stones[i] 放入容量为Sum/2的背包中，
//* 使得背包中的stones的重量和最接近Sum/2。由于是求剩下的最小的值，故最小值结果则为：abs(Sum - f[n][Sum/2] - f[n][Sum/2])
func lastStoneWeightII(stones []int) int {
	sum, n := 0, len(stones)
	for i := 0; i < n; i++ {
		sum += stones[i]
	}
	t := sum / 2
	f := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]int, t+1)
	}

	for i := 1; i < n+1; i++ {
		for j := 0; j <= t; j++ {
			if j < stones[i-1] {
				f[i][j] = f[i-1][j]
				continue
			}
			f[i][j] = max(f[i-1][j], f[i-1][j-stones[i-1]]+stones[i-1])
		}
	}
	return abs(sum - 2*f[n][t])
}

//* 一维数组优化空间。
//TODO: 由于 f[i][j] 状态值依赖于  f[i-1][j] 的状态值,可将其降为1维数组遍历，
//TODO: 其次，由于其需要获取上一次的 f[i-1][j-stones[i-1]] 的值，故应该从高位依次更新至低位，免得覆盖了f[i-1]的值。
func lastStoneWeightII_1(stones []int) int {
	n := len(stones)
	sum := 0
	for i := 0; i < n; i++ {
		sum += stones[i]
	}
	t := sum / 2
	f := make([]int, t+1)

	for i := 1; i <= n; i++ {
		x := stones[i-1]
		for j := t; j >= x; j-- {
			f[j] = max(f[j], f[j-x]+x)
		}
	}
	return abs(sum - f[t] - f[t])
}

// func abs(a int) int {
// 	if a >= 0 {
// 		return a
// 	}
// 	return -1 * a
// }
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
