package leetcode

import "container/heap"

/*
* 215. 数组中的第K个最大元素
* 在未排序的数组中找到第 k 个最大的元素。
* 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

! 说明:
* 你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/kth-largest-element-in-an-array/solution/shu-zu-zhong-de-di-kge-zui-da-yuan-su-by-leetcode-/
*/

//堆排实现
func findKthLargest(nums []int, k int) int {
	heapSize := len(nums)
	buildMaxHeap1(nums, heapSize)
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify1(nums, 0, heapSize)
	}
	return nums[0]
}

func buildMaxHeap1(a []int, heapSize int) {
	for i := heapSize / 2; i >= 0; i-- {
		maxHeapify1(a, i, heapSize)
	}
}

func maxHeapify1(a []int, i, heapSize int) {
	l, r, largest := i*2+1, i*2+2, i
	if l < heapSize && a[l] > a[largest] {
		largest = l
	}
	if r < heapSize && a[r] > a[largest] {
		largest = r
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapify1(a, largest, heapSize)
	}
}

//利用golang自身heap包的实现
type maxHeap []int

func (n maxHeap) Len() int           { return len(n) }
func (n maxHeap) Less(i, j int) bool { return n[i] > n[j] }
func (n maxHeap) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n *maxHeap) Push(x interface{}) {
	*n = append(*n, x.(int))
}
func (n *maxHeap) Pop() interface{} {
	old := *n
	nLen := len(old)
	x := old[nLen-1]
	*n = old[:nLen-1]
	return x
}
func findKthLargest1(nums []int, k int) int {
	var maxnums maxHeap = nums
	heap.Init(&maxnums)
	for i := 0; i < k; i++ {
		max := heap.Pop(&maxnums)
		if i == k-1 {
			return max.(int)
		}
	}
	return 0
}
