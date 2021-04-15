package leetcode

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
