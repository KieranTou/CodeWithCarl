package main

type SingleNode struct {
	Val  int
	Next *SingleNode
}

type SingleLinkedList struct {
	dummyHead *SingleNode
	Size      int
}

func main() {

}

func (list *SingleLinkedList) reverseList(head *SingleNode) *SingleNode {
	var temp *SingleNode
	cur := head
	var pre *SingleNode = nil

	for cur != nil {
		temp = cur.Next
		cur.Next = pre

		pre = cur
		cur = temp
	}
	return pre
}

// 递归法
func (list *SingleLinkedList) reverseList1(head *SingleNode) *SingleNode {
	return help(nil, head)
}
func help(pre, head *SingleNode) *SingleNode {
	if head == nil {
		return pre
	}
	next := head.Next
	head.Next = pre
	return help(head, next)
}
