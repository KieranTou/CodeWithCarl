package main

import "fmt"

//mid := left + ((right - left) >> 1)

func main() {
	// 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target.
	// 如果目标值存在返回下标，否则返回 -1。

	// 满足 1有序数组 2无重复元素 可考虑二分法
	// 分为两种方法 [left,right] [left,right)

	// 二分法其实就是不断将区间缩小至一点的过程

	array := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0}
	// 虽说是array 但传入必须是slice 如果上面中括号里写了10或...(自动判断长度) 这里就无法传入函数
	res := search1(array, 7)
	fmt.Println(res)
}

// 第一种写法 左闭右闭
func search1(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	// go无while循环 用for
	// 因为两边都是闭区间 low==high有意义
	for low <= high {
		//这样写是为了防止写(low + high)/2 导致大数溢出
		mid := low + (high-low)/2
		if target < nums[mid] {
			high = mid - 1
		} else if target > nums[mid] {
			low = mid + 1
		} else {
			// 这里先判断nums[mid] == target 返回mid 更好看
			return mid
		}
	}
	return -1
}

// 第二种写法 左闭右开
func search2(nums []int, target int) int {
	low := 0
	// 因为右边是开区间取不到 所以这里high不需-1
	high := len(nums)
	// 这里
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if target > nums[mid] {
			low = mid + 1
		} else {
			// 因为右侧是开区间 如果用等于mid-1会导致漏值
			high = mid
		}
	}
	return -1
}
