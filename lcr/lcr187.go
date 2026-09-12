package lcr

// f(m,n)=(f(m-1,n)+n)%m
func iceBreakingGame(num int, target int) int {
	var dfs func(num int, target int) int

	dfs = func(num int, target int) int {
		if num == 1 {
			return 0
		}
		return (dfs(num-1, target) + target) % num
	}
	return dfs(num, target)
}

func iceBreakingGameII(num int, target int) int {
	// index: num
	// value: the last remaining position index
	ans := make([]int, num+1)

	ans[1] = 0
	for i := 2; i <= num; i++ {
		ans[i] = (ans[i-1] + target) % (i)
	}
	return ans[num] - 1

}
