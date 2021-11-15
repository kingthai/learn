package interview

import "fmt"

type ListNode struct {
	Next  *ListNode
	Value int
}

func Reverse(node *ListNode) *ListNode {
	if node.Next == nil {
		return node
	}

	last := Reverse(node.Next)
	node.Next.Next = node
	node.Next = nil
	return last
}


func TestZuiyou () {
	node := &ListNode{
		Next: nil,
		Value: 2,
	}

	head := &ListNode{}
	head.Value = 1
	head.Next = node

	res := Reverse(head)

	fmt.Printf("list: %+v", res)

}


func down(a []int, i, n int) {
	for {
		child := 2*i+1

		if child >= n {
			break
		}
		if child + 1 < n && a[child+1] < a[child] {
			child++
		}

		if a[i] < a[child] {
			break
		}

		a[i], a[child] = a[child], a[i]

		i = child

	}


}


//awk {$a[1]++}END {for }
//awk {}   |sort | uniq -c | sort -nr
