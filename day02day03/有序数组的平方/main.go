package main

import (
	"fmt"
	"sort"
)

// 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

type Slice []int

func (slice Slice) Len() int {
	return len(slice)
}
func (slice Slice) Less(i, j int) bool {
	return slice[i] < slice[j]
}
func (slice Slice) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func main() {
	arr := []int{-4, -1, 0, 3, 10}
	force(arr)
	res := sortedSquares(arr)
	fmt.Println(res)
}

//暴力解法 平方后排序
func force(nums []int) []int {
	fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		nums[i] *= nums[i]
	}
	sort.Sort(Slice(nums))
	fmt.Println(nums)
	return nums
}

func sortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	k := n - 1
	for i, j := 0, n-1; i <= j; {
		lm, lr := nums[i]*nums[i], nums[j]*nums[j]
		if lm > lr {
			res[k] = lm
			i++
		} else {
			res[k] = lr
			j--
		}
		k--
	}
	return res
}
