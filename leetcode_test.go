package leetcode_test

import (
	"testing"
)

func TestProject(t *testing.T) {
	// coins := []int{186, 419, 83, 408}
	// amount := 6249
	// ans := leetcode.CoinChange2(coins, amount)
	// fmt.Println(ans)
	nums := make([]int, 10000)
	for i := 10000; i > 0; i-- {
		nums[10000-i] = i
	}
	// timestart := time.Now()
	// nums = leetcode.MergeSort(nums)
	// timeSpend := time.Now().Sub(timestart)
	// fmt.Println(nums)
	// fmt.Println(timeSpend)
	// for len(nums) > 0 {
	// 	fmt.Println(len(nums))
	// 	nums = nums[1:]
	// 	fmt.Println(nums)
	// }

}
