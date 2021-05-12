package leetcode_test

import (
	"fmt"
	"leetcode"
	"testing"
)

func TestProject(t *testing.T) {
	// coins := []int{186, 419, 83, 408}
	// amount := 6249
	// ans := leetcode.CoinChange2(coins, amount)
	// fmt.Println(ans)
	nums := make([]int, 7)
	for i := 0; i < 7; i++ {
		nums[i] = i + 1
	}
	// timestart := time.Now()
	// nums = leetcode.MergeSort(nums)
	// timeSpend := time.Now().Sub(timestart)
	// fmt.Println(nums)
	// fmt.Println(timeSpend)
	// for len(nums) > 0 {
	// 	fmt.Println(len(nums))
	// 	nums = nums[1:]
	// fmt.Println(nums)
	// }
	// var nums1 leetcode.IntHeap = nums
	// heap.Init(&nums1)
	// nums = leetcode.HeapSort1(nums)
	// sort.Sort(nums)
	// timeSpend := time.Now().Sub(timestart)
	// max1 := heap.Pop(&nums)
	// res := max1.(int)
	// fmt.Println(res)
	// fmt.Println("Time Spend:", timeSpend, "\r\n", max1)
	// fmt.Println(nums)
	// var nums1 []int
	// nums1 = nums
	// fmt.Println(nums1)
	// nums := leetcode.Minnum{3, 30, 34, 5, 9}
	// nums1 := []int{3, 30, 34, 5, 9}
	// var n leetcode.Minnum = nums1
	// sort.Sort(n)
	// fmt.Println(n)

	// root := leetcode.BinaryTreeConstructor(nums)
	// leetcode.BinaryTreeTraverse(root)

	// leetcode.QuickSort(nums)
	// timeSpend := time.Now().Sub(timestart)
	// fmt.Println(nums)
	nums = []int{7, 2, 5, 3, 11, 9, 1}
	ans := leetcode.MonotoneStack(nums)
	fmt.Println(ans)

}
