package main

// 使用数组来做哈希的题目，是因为题目都限制了数值的大小
// 而且如果哈希值比较少、特别分散、跨度非常大，使用数组就造成空间的极大浪费
// 直接使用set 不仅占用空间比数组大，而且速度要比数组慢，set把数值映射到key上都要做hash计算的。

// go要用map模拟set 因为key是唯一值
// map的value值是布尔型，这会导致set多占用内存空间，解决这个问题，则可以将其替换为空结构。在Go中，空结构通常不使用任何内存。
func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]struct{}, 0)
	res := make([]int, 0)
	// 将第一个数组并入集合
	for _, v := range nums1 {
		if _, ok := set[v]; !ok {
			set[v] = struct{}{}
		}
	}
	// 看第二个数组元素是否存在
	for _, v := range nums2 {
		if _, ok := set[v]; ok {
			res = append(res, v)
			delete(set, v)
		}
	}
	return res
}
