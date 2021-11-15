package main

import (
	"fmt"
	"sort"
)

func main() {

	nums := []int{1, 2, 3, -3, -1}
	sort.Ints(nums)
	res := make([][]int, 0)
	getCollection(nums, &res)

	for i := range res {
		fmt.Println(res[i])
	}

}

type TreeNode struct {
	Val int
	LNode *TreeNode
	RNode *TreeNode
}
func getFather(root, node1, node2 *TreeNode) *TreeNode{
	if root == nil {
		return nil
	}

	if root.Val == node1.Val || root.Val == node2.Val {
		return root
	}

	l := getFather(root.LNode, node1, node2)
	r := getFather(root.RNode, node1, node2)

	if l != nil && r != nil {
		return root
	}
	if l == nil {
		return r
	}
	return l


}


func dfs (root, node *TreeNode) *TreeNode {

}


// res 存结果
func getCollection(nums []int, res *[][]int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= 0 {
			break
		}

		target := -nums[i]

		// 2数之和
		l, r := i+1, len(nums)-1
		for l<r {
			if nums[l] + nums[r] < target {
				l++
			}else if nums[l] + nums[r] > target {
				r--
			}else {
				//fmt.Println(nums[i], nums[l], nums[l])
				*res = append(*res, []int{nums[i], nums[l], nums[l]})
				l++
				r--
			}
		}
	}

}
