package leetcode

import (
	"container/heap"
	"math/rand"
	"time"
)

/*
* 冒泡排序（稳定排序算法）：内部排序
*	1.相邻的2各元素比较，大的向后移，经过一轮比较，做大的元素排在最后
*	2.第二轮，第二大的元素排倒数第二个位置
*	3.直到全部排好

* 时间复杂度：O(n^2)；
* 空间复杂度：O(1)
 */
func BubbleSort(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return nums
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i; j++ {
			if j+1 < len(nums) && nums[j] > nums[j+1] {
				nums[j+1], nums[j] = nums[j], nums[j+1]
			}
		}
	}
	return nums
}

/*
* 选择排序(不稳定)：内部排序
* 1. 首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
* 2. 再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
* 3. 重复第二步，直到所有元素均排序完毕。

! 本质和冒泡排序一样
* 时间复杂度：总是O(n^2)
* 空间复杂度：O(1)
*/

func SelectionSort(nums []int) []int {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

/*
* 插入排序（稳定）：内部排序
* 1. 将第一待排序序列第一个元素看做一个有序序列，把第二个元素到最后一个元素当成是未排序序列。
* 2. 从头到尾依次扫描未排序序列，将扫描到的每个元素插入有序序列的适当位置。
* （如果待插入的元素与有序序列中的某个元素相等，则将待插入元素插入到相等元素的后面。）

* 时间复杂度：O(n^2)；最好的情况（初始为顺序）O(n)，最差情况（初始为逆序）：O(n^2)
* 空间复杂度：O(1)
 */
func InsertionSort(nums []int) []int {
	for i := range nums {
		preIndex := i - 1
		current := nums[i]
		//逐个比较当前元素与当前的有序数组的元素，找到当前元素在有序数组的位置。
		//当前元素比有序数组的元素小，则preIndex-1，即向前位置的元素比较
		for preIndex >= 0 && nums[preIndex] > current {
			nums[preIndex+1] = nums[preIndex]
			preIndex -= 1
		}
		//然后更新当前元素在有序数组种的位置
		nums[preIndex+1] = current
	}
	return nums
}

/*
* 希尔排序(不稳定)：内部排序
* 1. 选择一个增量序列 t1，t2，……，tk，其中 ti > tj, tk = 1；
* 2. 按增量序列个数 k，对序列进行 k 趟排序；
* 3. 每趟排序，根据对应的增量 ti，将待排序列分割成若干长度为 m 的子序列，分别对各子表进行直接插入排序。
*	仅增量因子为 1 时，整个序列作为一个表来处理，表长度即为整个序列的长度。

* 希尔排序的基本思想是：
* 	先将整个待排序的记录序列分割成为若干子序列分别进行直接插入排序，
* 	待整个序列中的记录“基本有序”时，再对全体记录进行依次直接插入排序。

* 时间复杂度：O(nlogn)
* 空间复杂度：O(1)
 */

func ShellSort(nums []int) []int {
	length := len(nums)
	gap := 1
	for gap < len(nums)/3 {
		gap = gap*3 + 1
	}
	for gap > 0 {
		for i := gap; i < length; i++ {
			temp := nums[i]
			j := i - gap
			for j >= 0 && nums[j] > temp {
				nums[j+gap] = nums[j]
				j -= gap
			}
			nums[j+gap] = temp
		}
		gap = gap / 3
	}
	return nums
}

/*
! 归并排序（稳定）
* 1. 申请空间，使其大小为两个已经排序序列之和，该空间用来存放合并后的序列；
* 2. 设定两个指针，最初位置分别为两个已经排序序列的起始位置；
* 3. 比较两个指针所指向的元素，选择相对小的元素放入到合并空间，并移动指针到下一位置；
* 4. 重复步骤 3 直到某一指针达到序列尾；
* 5. 将另一序列剩下的所有元素直接复制到合并序列尾。
*/

//自上而下的递归实现：
func MergeSort(nums []int) []int {
	length := len(nums)
	if length < 2 {
		return nums
	}
	middle := length / 2
	left := nums[0:middle]
	right := nums[middle:]
	return merge(MergeSort(left), MergeSort(right))
}

func merge(left []int, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	//将剩余的左边部分加入结果
	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}
	//将剩余的右边部分加入结果
	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}

	return result
}

//自下而上的迭代实现

/*
! 快速排序
* 快速排序算法由 C. A. R. Hoare 在 1960 年提出。它的时间复杂度也是 O(nlogn)，
* 但它在时间复杂度为 O(nlogn) 级的几种排序算法中，大多数情况下效率更高，所以快速排序的应用非常广泛。
* 再加上快速排序所采用的分治思想非常实用，使得快速排序深受面试官的青睐，所以掌握快速排序的思想尤为重要。

! 快速排序算法的基本思想是：
* 	从数组中取出一个数，称之为基数（pivot）
* 	遍历数组，将比基数大的数字放到它的右边，比基数小的数字放到它的左边。遍历完成后，数组被分成了左右两个区域
* 	将左右两个区域视为两个数组，重复前两个步骤，直到排序完成

* 事实上，快速排序的每一次遍历，都将基数摆到了最终位置上。第一轮遍历排好 1 个基数，
* 第二轮遍历排好 2 个基数（每个区域一个基数，但如果某个区域为空，则此轮只能排好一个基数），
* 第三轮遍历排好 4 个基数（同理，最差的情况下，只能排好一个基数），以此类推。
* 总遍历次数为 logn～n 次，每轮遍历的时间复杂度为 O(n)，所以很容易分析出快速排序的时间复杂度为 O(nlogn)～ O(n^2),
* 平均时间复杂度为 O(nlogn)O(nlogn)。

* 空间复杂度与递归的层数有关，每层递归会生成一些临时变量，
* 所以空间复杂度为 O(logn) ~ O(n)，平均空间复杂度为 O(logn)。

! 基数的选择对于快排时间复杂度的影响：
! 基数的选择没有固定标准，随意选择区间内任何一个数字做基数都可以。通常来讲有三种选择方式：
*	选择第一个元素作为基数
*	选择最后一个元素作为基数
*	选择区间内一个随机元素作为基数
! 选择区间内一个随机元素作为基数的平均时间复杂度是最优的.

! 快速排序递归框架
func QuickSort(arr []int){
	quickSort(arr, 0, len(arr)-1)
}
func quickSort(arr []int, start, end int) {
    * 将数组分区，并获得中间值的下标
    middle := partition(arr, start, end);
    * 对左边区域快速排序
    quickSort(arr, start, middle-1);
    * 对右边区域快速排序
    quickSort(arr, middle+1, end);
}
public static int partition(arr []int, start, end int) {
    TODO: 将 arr 从 start 到 end 分区，左边区域比基数小，右边区域比基数大，然后返回中间值的下标
}

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/sort-algorithms/eul7hm/
*/

func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, start, end int) {
	if end-start <= 0 {
		return
	}
	mid := partition(nums, start, end)
	quickSort(nums, start, mid-1)
	quickSort(nums, mid+1, end)
}

func partition(nums []int, start, end int) int {
	//TODO:将 arr 从 start 到 end 分区，左边区域比基数小，右边区域比基数大，然后返回中间值的下标
	//选择基数 (随机选择)
	rand.Seed(time.Now().Unix())
	randPivot := start + rand.Intn(end-start)
	//中间选择
	// randPivot := (start + end) / 2
	//TODO: 将选定的比较基数，放置最后边。（便于遍历区间比较大小，找到基数真实位置）
	nums[randPivot], nums[end] = nums[end], nums[randPivot]
	i := start - 1
	//TODO: 遍历区间，如果比选定的基数小，则放置在基数的前面位置
	for j := start; j < end; j++ {
		if nums[j] < nums[end] {
			i += 1
			nums[j], nums[i] = nums[i], nums[j]
		}
	}
	//TODO：将选定的基数放回其正确的位置
	i += 1
	nums[i], nums[end] = nums[end], nums[i]
	return i
}

/*
* 堆排序
* 堆排序框架：
func HeapSort(nums []int) []int {
	heapSize := len(nums)
	//*建大根堆
	buildMaxHeap(nums, heapSize)
	//* 将堆顶的最大元素放置在队尾，并更新堆的大小（即，堆大小减1）。
	//* 直到堆大小归 0.
	for i := len(nums) - 1; i >= 1; i-- {
		交换位置
		//* 对子树大根堆化。
		maxHeapify(nums, 0, heapSize)
	}
	return nums
}
func buildMaxHeap(a []int, heapSize int) {
}

func maxHeapify(a []int, i, heapSize int) {
}

*/

//*堆排序 手撸实现
func HeapSort(nums []int) []int {
	heapSize := len(nums)
	//* 建大根堆
	buildMaxHeap(nums, heapSize)
	//* 将堆顶的最大元素放置在队尾，并更新堆的大小（即，堆大小减1）。
	//* 直到堆大小归 0.
	for i := len(nums) - 1; i >= 1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		//* 堆更新的堆，重新大根堆化
		maxHeapify(nums, 0, heapSize)
	}
	return nums
}

func buildMaxHeap(a []int, heapSize int) {
	//* 由下自上进行大根堆化，即，从最后的子树进行大根堆化。
	//* 每次将子树中最大的元素放置在根部。
	for i := heapSize / 2; i >= 0; i-- {
		maxHeapify(a, i, heapSize)
	}
}

func maxHeapify(a []int, i, heapSize int) {
	l, r, largest := i*2+1, i*2+2, i
	if l < heapSize && a[l] > a[largest] {
		largest = l
	}
	if r < heapSize && a[r] > a[largest] {
		largest = r
	}
	if largest != i { //*说明原来的最大数变了
		//*更新子树根为最大数
		a[i], a[largest] = a[largest], a[i]
		//*向下调整，看看调整后的以larges为的根的子树是否符合大根化。
		maxHeapify(a, largest, heapSize)
	}
}

//golang中实现小根堆
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } //i<j时为小根堆，i>j时为大根堆
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
func HeapSort1(nums []int) []int {
	var nums1 IntHeap = nums
	var result []int
	heap.Init(&nums1)
	for i := 0; i < len(nums); i++ {
		n := heap.Pop(&nums1).(int)
		result = append(result, n)
	}
	return result
}
