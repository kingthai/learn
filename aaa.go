package main

func main() {

}

type Node struct {
	Val int
	LC *Node
	RC *Node
}

func tree(root *Node) bool {
	if root == nil {
		return false
	}

	// bfs
	queue := make([]*Node, 0)
	queue = append(queue, root)

	for len(queue) > 0 {

		l, r := 0, len(queue)-1
		for l < r {
			if queue[l].Val != queue[r].Val {
				return false
			}
			l++
			r--
		}

		// 把子节点 添加到队列
		n := len(queue)
		newQ := make([]*Node, 0)
		for n > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.LC != nil {
				newQ = append(newQ, node.LC)
			}
			if node.RC != nil {
				newQ = append(newQ, node.RC)
			}

			if node.LC != nil && node.RC == nil {
				return false
			}else if node.LC == nil && node.RC != nil {
				return false
			}
		}
		queue = newQ

	}
	return true
}