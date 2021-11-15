package alth

//给定一个含不同整数的集合，返回其所有的子集。
//输入：[1,2,3]
//输出：
//[
//[3],
//[1],
//[2],
//[1,2,3],
//[1,3],
//[2,3],
//[1,2],
//[]
//]

//func main() {
	//choice := make([]int, 0)
	//res := make([][]int, 0)
	//
	//nums := []int{1,2,3}
	//getCollect(nums, choice, &res)
	//fmt.Printf("res: %+v", res)
//}

type Node struct {
	Left, Right *Node
}


func getCollect(nums, choice []int, res *[][]int) {
	// 子集添加
	*res = append(*res, append([]int{}, choice...))

	for i := 0; i < len(nums); i++ {
		choice = append(choice, nums[i])
		//fmt.Println(choice)
		getCollect(nums[:i], choice, res)
		choice = choice[:len(choice)-1]
	}
}
