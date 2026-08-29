package lcr

func trainWays(num int) int {

	if num == 0 {
		return 1
	}
	if num == 1 {
		return 1
	}

	ans := make([]int, num+1)
	ans[0] = 1
	ans[1] = 1

	for i := 2; i < num+1; i++ {
		ans[i] = (ans[i-1] + ans[i-2]) % 1000000007
	}
	return ans[num]
}
