package middle

func punishmentNumber(n int) int {
	result := make([]int, 0)

	for i := 1; i <= n; i++ {
		if validatePunishNum(i*i, i) {
			result = append(result, i)
		}

	}

	sum := 0

	for _, v := range result {
		sum += v * v
	}

	return sum

}

func validatePunishNum(target int, num int) bool {
	if target < num {
		return false
	}
	if target == num {
		return true
	}

	i := 10

	for {
		if i > target {
			return false
		}
		if validatePunishNum(target/i, num-target%i) {
			return true
		}

		i *= 10
	}

}
