package alth

import "math"

// 从上到下递归计算
func coinChange(coins []int, amount int) int {
    capT := make(map[int]int)
    a := helper(coins, amount, capT)
    return a
}

func helper (coins []int, amount int, capT map[int]int) int {
	if amount == 0 {
		return 0
	}else if amount < 0 {
		return -1
	}

	if v, ok := capT[amount]; ok {
		return v
	}

	res := math.MaxInt32
	for _, c := range coins {
		sub := helper(coins, amount - c, capT)
		if sub == -1 {
			continue // 无解
		}

		if sub < res {
			res = sub + 1
		}
	}

	if res != math.MaxInt32 {
		capT[amount] = res
	}else {
		capT[amount] =  -1
	}
	return capT[amount]
}
