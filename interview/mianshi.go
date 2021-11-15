package interview

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		// 当前值 vs dp[i-1]+当前值
		dp[i] = max(nums[i], nums[i]+dp[i-1])
	}

	// 获取最大的集合
	res := math.MinInt32
	for i := 0; i < len(nums); i++ {
		res = max(res, dp[i])
	}
	return res
}
func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func main() {
	nums := []int{1, 3, -5, 8}
	res := maxSubArray(nums)
	fmt.Println(res)
}
