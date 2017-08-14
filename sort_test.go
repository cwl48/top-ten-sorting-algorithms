package sort

import (
	"fmt"
	"testing"
)

var nums = []int{445, 1, 2, 6, 66, 8, 9, 8, 5, 5, 3, 5, 6, 846, 55}

func Test_bubbleSort(t *testing.T) {

	result := bubbleSort(nums)
	n, s := checkSort(result)

	if n == 0 {
		t.Error("结果为无序的:", result)
	} else {
		fmt.Println("结果为", s, ":", result)
	}
}

func Test_quickSort(t *testing.T) {

	result := quickSort(nums)
	n, s := checkSort(result)

	if n == 0 {
		t.Error("结果为无序的:", result)
	} else {
		fmt.Println("结果为", s, ":", result)
	}
}

func Test_selectSort(t *testing.T) {

	result := selectSort(nums)
	n, s := checkSort(result)

	if n == 0 {
		t.Error("结果为无序的:", result)
	} else {
		fmt.Println("结果为", s, ":", result)
	}
}

func Test_insertSort(t *testing.T) {

	result := insertSort(nums)

	n, s := checkSort(result)

	if n == 0 {
		t.Error("结果为无序的:", result)
	} else {
		fmt.Println("结果为", s, ":", result)
	}
}

func Test_shellSort(t *testing.T) {

	result := shellSort(nums)
	n, s := checkSort(result)

	if n == 0 {
		t.Error("结果为无序的:", result)
	} else {
		fmt.Println("结果为", s, ":", result)
	}
}

func Test_mergeSort(t *testing.T) {

	result := mergeSort(nums)
	n, s := checkSort(result)

	if n == 0 {
		t.Error("结果为无序的:", result)
	} else {
		fmt.Println("结果为", s, ":", result)
	}
}

func Test_heapSort(t *testing.T) {

	result := heapSort(nums)
	n, s := checkSort(result)

	if n == 0 {
		t.Error("结果为无序的:", result)
	} else {
		fmt.Println("结果为", s, ":", result)
	}
}

func Test_countSort(t *testing.T) {
	result := countSort(nums)
	n, s := checkSort(result)

	if n == 0 {
		t.Error("结果为无序的:", result)
	} else {
		fmt.Println("结果为", s, ":", result)
	}
}

func Test_bucketSort(t *testing.T) {
	result := bucketSort(nums)
	n, s := checkSort(result)

	if n == 0 {
		t.Error("结果为无序的:", result)
	} else {
		fmt.Println("结果为", s, ":", result)
	}
}
