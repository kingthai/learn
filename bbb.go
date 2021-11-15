package main

import "fmt"

// m*n 2维数组 顺时针输出数组   空间复杂度O（1）

// id uid other_uid c_time u_time  --- 表 user_rela
// uid status  c_time u_time 用户状态 是否正常 0-正常 1-异常  -- 表 user
// 查询 用户正常好友列表
// join 变为   status冗余
// 冗余操作  join去掉

// id uid other_uid other_user_status
// 步骤:
// 添加other_user_status 列

// 先上线同步代码
//select status from user
//insert into user_rela (id,uid,other_uid,other_user_status,c_time,u_time) values ();



// 同步status到 user_rela
// update user_rela as a left join user as b on a.other_uid = b.uid set a.other_user_status = b.status
// where a.u_time > xxxx and b.u_time > xxxx

// 查询语句修改
// select other_uid from user_rela where other_user_status = 0 and uid = x

func main () {
	nums := make([][]int, 0)
	nums = append(nums, []int{1,2,3})
	nums = append(nums, []int{4,5,6})
	nums = append(nums, []int{7,8,9})
	//nums = append(nums, []int{1})
	//nums = append(nums, []int{4})
	//nums = append(nums, []int{7})

	res := printArr(nums)
	fmt.Println(res)

}

func printArr(nums [][]int) []int {
	if len(nums) == 0 {
		return nil
	}

	r, c := len(nums), len(nums[0])

	if c ==0 {
		return nil
	}

	// 定义返回结果
	res := make([]int, 0)

	// 只有一列 特殊处理
	if c == 1 {
		for i:=0;i<r;i++ {
			res = append(res, nums[i][0])
		}
		return res
	}


	// 起始位置
	start := 0

	for r > start * 2 && c > start * 2 {
		endX := c-1-start
		endY := r-1-start


		// 从左右到右打印 第一行的边
		for i:=start;i<=endX;i++ {
			res = append(res, nums[start][i])
		}

		// 从最右到最下面  最右边的边
		if start < endY {
			for i:=start+1;i<=endY;i++ {
				res = append(res, nums[i][endX])
			}
		}

		// 从最右到左打印 最下面的边
		if start < endX && start < endY {
			for i:=endX-1;i>=start;i-- {
				res = append(res, nums[endY][i])
			}
		}

		// 从最左到最上
		if start < endY - 1 {
			for i := endY-1;i>=start+1;i-- {
				res = append(res, nums[i][start])
			}
		}
		start++
	}
	return res
}