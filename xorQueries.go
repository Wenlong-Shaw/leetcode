package leetcode

/*
 */

//暴力法. 遍历queries的每个查询区间数组，然后求解result[i]。
//存在问题：会重复求解重合的区间段。
//解决思路：借鉴dp中的状态数组，将求得的每个区间保留
func xorQueries(arr []int, queries [][]int) []int {
	result := make([]int, 0)
	a := 0
	for i := 0; i < len(queries); i++ {
		for j := queries[i][0]; j <= queries[i][1]; j++ {
			a ^= arr[j]
		}
		result = append(result, a)
	}
	return result
}

// 状态（结果）数组保存法
func xorQueries1(arr []int, queries [][]int) []int {
	xorState := make([]int, len(arr)+1)

	for i := 0; i < len(arr); i++ {
		xorState[i+1] = xorState[i] ^ arr[i]
	}
	ans := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		ans[i] = xorState[queries[i][0]] ^ xorState[queries[i][1]+1] //异或首、尾区间等于中间区间异或值
	}
	return ans
}
