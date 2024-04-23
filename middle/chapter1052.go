package middle

func MaxSatisfied(customers []int, grumpy []int, minutes int) int {
	m := len(customers)
	sum := 0
	for i := 0; i < m; i++ {
		if grumpy[i] == 0 {
			sum += customers[i]
		}
	}

	initSum := 0
	for i := 0; i <= (minutes - 1); i++ {
		if grumpy[i] == 1 {
			initSum += customers[i]
		}
	}

	left := 0
	//右端点 ： left+minutes-1
	dp := make([]int, m-(left+minutes-1))
	dp[0] = initSum

	left++
	for left+minutes-1 < m {
		if grumpy[left-1] == 1 {
			initSum -= customers[left-1]
		}
		if grumpy[left+minutes-1] == 1 {
			initSum += customers[left+minutes-1]
		}
		dp[left] = initSum
		left++
	}

	max := 0
	for i := 0; i < len(dp); i++ {
		if max < dp[i] {
			max = dp[i]
		}

	}
	return max + sum

}
