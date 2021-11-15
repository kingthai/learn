package alth

type ListNode struct {
	     Val int
	     Next *ListNode
}

// 迭代
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	var prev *ListNode
	cur, next := head, head

	for cur != nil {
		next = cur.Next
		cur.Next = prev
		// 更新指针位置
		prev = cur
		cur = next
	}

	return prev
}
