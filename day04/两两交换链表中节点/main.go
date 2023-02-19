package main

import ()

/**
			标准力扣定义
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	cur := dummy // 为什么要有cur 因为最后要返回头节点 对dummy操作的话头节点就找不到了
	// 把虚拟头节点当成 -1 节点就不会迷茫了
	for cur.Next != nil && cur.Next.Next != nil {
		temp := cur.Next
		temp1 := cur.Next.Next.Next

		cur.Next = cur.Next.Next // ? dummy.N = dummy.N.N ?
		cur.Next.Next = temp
		cur.Next.Next.Next = temp1
		cur = cur.Next.Next //自此cur与dummy无关 但之前已经将dummy指向了2

	}
	return dummy.Next

}
