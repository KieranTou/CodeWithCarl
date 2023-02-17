package main

import "fmt"

// 给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。
// 示例：
// 输入：s = 7, nums = [2,3,1,2,4,3] 输出：2 解释：子数组 [4,3] 是该条件下的长度最小的子数组。
func main() {
	// 其实就是双指针 条件不同的快慢双指针
	arr := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(arr)
	fmt.Println(minSubArrayLen(arr, 6))
}

func minSubArrayLen(nums []int, target int) int {
	i := 0
	l := len(nums)
	sum := 0
	res := l + 1
	for j := 0; j < l; j++ {
		sum += nums[j]
		for sum > target {
			subLength := j - i + 1
			if subLength < res {
				res = subLength
			}
			sum -= nums[i]
			i++
		}
	}
	if res == l+1 {
		return 0
	} else {
		return res
	}

}
