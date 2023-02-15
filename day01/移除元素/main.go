package main

import "fmt"

func main() {
	// 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
	// 不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并原地修改输入数组。
	// 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

	array := []int{3, 5, 7, 9, 2, 2, 6, 7, 2, 1, 0, 8, 4, 4}
	res := removeElementBeSide(array, 2)
	fmt.Println(res)
}

// 暴力解法 遍历数组 遇到就删掉然后把后面的元素均向前移动一个
func removeForce(nums []int, val int) int {
	length := len(nums)
	for i := 0; i < length; i++ {
		if nums[i] == val {
			for j := i; j < length-1; j++ {
				nums[j] = nums[j+1]
			}
			length--
		}
	}
	fmt.Println(nums)
	return length
}

// 我发现当有连续的2出现时 后面的一个并没有被删除掉 因为i没有后退 所以会略过值
// 记得 数组长度减少时 令i也后退一下 更改后---
func removeForceV2(nums []int, val int) int {
	length := len(nums)
	for i := 0; i < length; i++ {
		if nums[i] == val {
			for j := i; j < length-1; j++ {
				nums[j] = nums[j+1]
			}
			length--
			i--
		}
	}
	return length
}

// 双指针法（快慢指针法） 在数组、链表、字符串中很常见
func removeElement(nums []int, val int) int {
	length := len(nums)
	// res作为慢指针检索val i作为快指针向前走
	res := 0
	for i := 0; i < length; i++ {
		if nums[i] != val {
			nums[res] = nums[i]
			res++
		}
		// 当慢指针指到val时慢指针不用做任何操作 只需要将快指针前移一下即可
		// 这里快指针一直在前移 所以无需操作
	}
	// 好家伙 不光快 array后面的垃圾信息也可排除
	nums = nums[:res]
	fmt.Println(nums)
	return res
}

// 整点高端的 补充Go的相向双指针法
func removeElementBeSide(nums []int, val int) int {
	// 有点像二分查找的左闭右闭区间 所以下面是<=
	left := 0
	right := len(nums) - 1
	for left <= right {
		// 不断寻找左侧的val和右侧的非val 找到时交换位置 目的是将val全覆盖掉
		for left <= right && nums[left] != val {
			left++
		}
		for left <= right && nums[right] == val {
			right--
		}
		//各自找到后开始覆盖
		if left < right {
			nums[left] = nums[right]
			left++
			right--
		}
	}
	fmt.Println(nums)
	return left
}
