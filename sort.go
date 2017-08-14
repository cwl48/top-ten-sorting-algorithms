package sort

import (
	"fmt"
)

/*
	冒泡排序
*/
func bubbleSort(nums []int) []int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

/*
	快速排序
*/
func quickSort(nums []int) []int {

	if len(nums) <= 0 {
		return nums
	}

	mid, i := nums[0], 1
	l, h := 0, len(nums)-1

	for l < h {
		if nums[i] > mid {
			nums[i], nums[h] = nums[h], nums[i]
			h--
		} else {
			nums[i], nums[l] = nums[l], nums[i]
			l++
			i++
		}
	}
	nums[l] = mid
	quickSort(nums[:l])
	quickSort(nums[l+1:])
	return nums
}

/*
	选择排序
*/
func selectSort(nums []int) []int {

	//最小数的下标
	var minIndex int

	for i := 0; i < len(nums); i++ {

		minIndex = i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			nums[minIndex], nums[i] = nums[i], nums[minIndex]
		}
	}

	return nums
}

/*
	插入排序
*/
func insertSort(nums []int) []int {

	var temp int
	fmt.Println(nums)
	for i := 1; i < len(nums); i++ {

		// 需要往前插入的元素

		temp = nums[i]

		if temp < nums[i-1] {
			for j := i - 1; j >= 0; j-- {
				if nums[j] > temp {
					nums[j], nums[j+1] = nums[j+1], nums[j]
				}
			}
		}

	}
	return nums
}

/*
	希尔排序

*/
func shellSort(nums []int) []int {

	// 下标相隔t的为一组
	t, j, temp := len(nums)/2, 0, 0

	for t >= 1 {
		for i := t; i < len(nums); i++ {
			j = i - t
			temp = nums[i]
			for j >= 0 && temp < nums[j] {
				nums[j+t], nums[j] = nums[j], nums[j+t]
				j = j - t
			}
		}
		t = t / 2
	}

	return nums
}

/*
	归并排序
*/
func mergeSort(r []int) []int {
	length := len(r)
	if length <= 1 {
		return r
	}
	num := length / 2
	left := mergeSort(r[:num])
	right := mergeSort(r[num:])
	return merge(left, right)
}
func merge(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}

/*
	堆排序
*/
func heapSort(nums []int) []int {

	// 非叶子节点数数组总索引 从0开始
	n := len(nums) - 1

	// 先构造大顶堆
	for i := n / 2; i >= 0; i-- {
		initHeap(nums, i, n)
	}

	// 把顶节点的值和尾节点进行交换
	// 因为顶节点已经最大，所以交换后尾节点就为有序
	for i := len(nums) - 1; i >= 1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		n--
		initHeap(nums, 0, n)
	}

	return nums
}

//构造堆
func initHeap(nums []int, k int, n int) {

	for {
		i := 2 * k
		if i > n {
			break
		}

		if i < n && nums[i] < nums[i+1] {
			i++
		}
		if nums[k] >= nums[i] {
			break
		}
		nums[i], nums[k] = nums[k], nums[i]
		k = i
	}

}

/*
	计数排序
*/
func countSort(nums []int) []int {

	// 获取数组中的最大值
	max := getNumMax(nums)

	// 临时存储区
	// 临时区保存着nums数组值大小范围内所有的整数(每个整数对应一个数组下标)
	var temp = make([]int, max+1)

	//  把所有元素放入临时区中, 即把对应元素个数加一 即：[1,3,2]=>[0,1,1,1]
	for i := 0; i < len(nums); i++ {
		temp[nums[i]]++
	}

	// 改变temp存储为每个数前面已经存在的数的数量+自身得数量 即 ： [0,1,1,1]=>[0,1,2,3]
	// "相加自身前面数字的数量"
	for i := 1; i <= max; i++ {
		temp[i] = temp[i] + temp[i-1]
	}

	// 输出结果
	outNums := make([]int, len(nums))

	for j := len(nums) - 1; j >= 0; j-- {
		outNums[temp[nums[j]]-1] = nums[j]
		temp[nums[j]]-- // 对应前面的数--
	}

	return outNums
}

/*
	桶排序  简单版计数排序，省去 "相加自身前面数字的数量"

*/
func bucketSort(nums []int) []int {

	// 获取数组中的最大值
	max := getNumMax(nums)

	// 临时存储区
	// 临时区保存着nums数组值大小范围内所有的整数   每个整数对应一个数组下标
	var temp = make([]int, max+1)

	// 把数组中每个数字的出现的个数 赋值给对应的 temp[nums[i]] 即：[1,3,2]=>[0,1,1,1]
	for i := 0; i < len(nums); i++ {
		temp[nums[i]]++
	}

	// 输出结果
	outNums := make([]int, 0, max)

	// 与计数排序不同 对数值数量>0的 直接进行顺序输出
	for j := 0; j <= max; j++ {

		if temp[j] > 0 {
			for x := 0; x < temp[j]; x++ {
				outNums = append(outNums, j)
			}
		}

	}

	return outNums
}

/*
	获取数组的最大值
	@return 数组中的最大值
*/
func getNumMax(nums []int) int {
	var max int
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

// 判断是否有序 包括升序 降序
// @return 1.升序 2.降序 0.无序
func checkSort(nums []int) (int, string) {

	//是否有序
	orderly := true

	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			orderly = false
		}
	}

	if orderly {
		return 1, "升序"
	}

	orderly = true

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			orderly = false
		}
	}

	if orderly {
		return 2, "降序"
	}

	return 0, "无序"

}
