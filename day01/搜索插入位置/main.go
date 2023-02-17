package main

import "fmt"

func main() {
	// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
	// 你可以假设数组中无重复元素。
	arr := []int{2, 4, 6, 8, 9, 17, 23}
	fmt.Println(arr)
}

// 暴力解法
func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}

//二分法 排序数组 无重复元素 所以使用二分法
//二分法很看边界条件 习惯使用左闭右闭 所以这里只写一种
func searchInsert2(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return len(nums)
}
