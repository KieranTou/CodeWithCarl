// 在排序数组中查找元素的第一个和最后一个位置
// 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
// 如果数组中不存在目标值 target，返回 [-1, -1]。

package main

import "fmt"

func main() {
	arr := []int{2, 5, 7, 9, 9, 12, 16, 89}

	fmt.Println(arr)
}

// 略
func searchRange(nums []int, target int) []int {
	return []int{}
}

func searchLeft(nums []int, target int) int {
	left, right := 0, len(nums)
	border := -1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
			border = right
		}
	}
	return border
}

func searchRight(nums []int, target int) int {
	left, right := 0, len(nums)
	border := -1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
			border = left
		}
	}
	return border
}
