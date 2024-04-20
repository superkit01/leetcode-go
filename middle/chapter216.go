package middle

//dfs

func combinationSum3(k int, n int) [][]int {
	ans := make([][]int, 0)
	dfs216(1, n, make([]int, 0), &ans, k)
	return ans

}

func dfs216(index int, remaining int, current []int, ans *[][]int, k int) {
	if remaining < 0 {
		return
	}
	if remaining == 0 {
		if len(current) == k {
			temp := make([]int, 0)
			temp = append(temp, current...)
			*ans = append(*ans, temp)
			return
		}
		return
	}

	if index > 9 {
		return
	}
	//使用当前位置
	temp1 := make([]int, 0)
	temp1 = append(temp1, current...)
	temp1 = append(temp1, index)
	dfs216(index+1, remaining-index, temp1, ans, k)

	//不使用当前位置
	temp2 := make([]int, 0)
	temp2 = append(temp2, current...)
	dfs216(index+1, remaining, temp2, ans, k)
}
