package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

// 传入的是头节点 只要传进来头节点 整个链表即可得到
//虚拟头节点写法
func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{} //地址 即指针类型变量
	dummyHead.Next = head
	cur := dummyHead
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}

//常规写法
func removeElement2(head *ListNode, val int) *ListNode {
	// 下面这里不能是if 为什么？ 因为可能有连续多个值都为val 删到不是头节点即可
	for head != nil && head.Val == val {
		head = head.Next
		return head
	}
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
