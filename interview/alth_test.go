package interview

import (
	"fmt"
	"testing"
)

func TestArrangement(t *testing.T) {
	arr := []int{6,2,8,0,4,7,9,0,0,3,5}
	root := InitTree(arr...)
	//levelOrder(root)
	//findAllPaths(root)
	ret := findPath(root, 7, 6)
	for i := range ret {
		fmt.Println(ret[i].Val)
	}
}

func TestCoinChange(t *testing.T)  {
	coins := []int{1,2,5}
	res := coinChange(coins, 13)
	fmt.Println(res)
}


