package interview

import (
	"fmt"
	"learn/utils/intutil"
	"learn/utils/sliceutil"
)

var res [][]int

// 全排列
func arrangement(arr []int, track []int) {
	if len(track) == len(arr) {
		res = append(res, append([]int{}, track...))
		return
	}

	for i := 0; i < len(arr); i++ {
		var f bool
		for j := range track {
			if track[j] == arr[i] {
				f = true
				break
			}
		}
		if f {
			continue
		}
		track = append(track, arr[i])
		arrangement(arr, track)
		track = track[:len(track)-1]
	}
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var DefaultValue = -1

func InsertNodeToTree(tree *TreeNode, node *TreeNode) {
	if tree == nil {
		return
	}
	if tree.Val == DefaultValue {
		tree.Val = node.Val
		return
	}
	if node.Val > tree.Val {
		if tree.Right == nil {
			tree.Right = &TreeNode{Val: DefaultValue}
		}
		InsertNodeToTree(tree.Right, node)
	}
	if node.Val < tree.Val {
		if tree.Left == nil {
			tree.Left = &TreeNode{Val: DefaultValue}
		}
		InsertNodeToTree(tree.Left, node)
	}
}

func InitTree(values ...int) (root *TreeNode) {
	rootNode := TreeNode{Val: DefaultValue, Right: nil, Left: nil}
	for _, value := range values {
		node := TreeNode{Val: value}
		InsertNodeToTree(&rootNode, &node)
	}
	return &rootNode
}

// 二叉树层序遍历
func levelOrder(root *TreeNode) {
	ret := make([][]int, 0)
	fc := func(root *TreeNode) {
		stack := make([]*TreeNode, 0)
		stack = append(stack, root)

		for len(stack) > 0 {
			newStack := make([]*TreeNode, 0)
			tmp := make([]int, 0)
			for i := range stack {
				node := stack[i]
				if node.Left != nil {
					tmp = append(tmp, node.Left.Val)
					newStack = append(newStack, node.Left)
				}
				if node.Right != nil {
					tmp = append(tmp, node.Right.Val)
					newStack = append(newStack, node.Right)
				}
			}
			stack = newStack
			if len(tmp) > 0 {
				ret = append(ret, append([]int{}, tmp...))
			}
		}
	}

	ret = append(ret, append([]int{}, root.Val))
	fc(root)
	fmt.Println(ret)
}

// Given a binary tree, return all root-to-leaf paths.
// 找出二叉树 所有从根节点到叶子节点的路径
func findAllPaths(root *TreeNode) {
	ret := make([]string, 0)
	_findAllPaths(root, fmt.Sprintf("%d", root.Val), &ret)
	fmt.Println(ret)
}

func _findAllPaths(node *TreeNode, path string, ret *[]string) {
	if node.Left == nil && node.Right == nil {
		*ret = append(*ret, path)
	}
	if node.Left != nil {
		_findAllPaths(node.Left, path+fmt.Sprintf("->%d", node.Left.Val), ret)
	}
	if node.Right != nil {
		_findAllPaths(node.Right, path+fmt.Sprintf("->%d", node.Right.Val), ret)
	}
}

// 找最近公共祖先
func lca(root *TreeNode, p, q int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p || root.Val == q {
		return root
	}
	left := lca(root.Left, p, q)
	right := lca(root.Right, p, q)
	if left != nil && right != nil {
		return root
	} else if left != nil {
		return left
	}
	return right
}

// 给定一棵二叉树的根节点和两个任意节点，返回这两个节点之间的最短路径
func findPath(root *TreeNode, p, q int) []*TreeNode {
	if root == nil {
		return nil
	}

	left := findPath(root.Left, p, q)
	right := findPath(root.Right, p, q)
	if left == nil && right == nil {
		if root.Val == p || root.Val == q {
			ret := make([]*TreeNode, 0)
			ret = append(ret, root)
			return ret
		} else {
			return nil // 没找到公共祖先
		}
	} else if left != nil && right != nil {
		ret := make([]*TreeNode, 0)
		ret = append(ret, left...)
		ret = append(ret, root)
		sliceutil.ReverseAny(right)
		ret = append(ret, right...)
		return ret
	} else if left != nil {
		ret := make([]*TreeNode, 0)
		ret = append(ret, root)
		ret = append(ret, left...)
		return ret

	} else {
		ret := make([]*TreeNode, 0)
		ret = append(ret, root)
		ret = append(ret, right...)
		return ret
	}
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 0; i < amount+1; i++ {
		for j := range coins {
			if i-coins[j] < 0 {
				continue
			}
			dp[i] = intutil.Min(dp[i], 1+dp[i-coins[j]])
		}
	}
	fmt.Println(dp)
	if dp[amount] == amount+1 {
		return amount
	}
	return dp[amount]
}
