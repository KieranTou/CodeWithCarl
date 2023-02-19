package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := &ListNode{Next: head}, &ListNode{Next: head}
	for ; n > 0; n-- {
		fast = fast.Next
	}
	//fast走了n步 此时若走完slow是倒数n+1个
	// 让slow是倒数n-1才能操作
	fast = fast.Next.Next
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
