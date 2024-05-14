package middle

func minimumRounds(tasks []int) int {
	cnt := make(map[int]int, 0)
	for _, v := range tasks {
		cnt[v]++
	}

	ans := 0
	for _, v := range cnt {
		if v == 1 {
			return -1
		}
		if v%3 != 0 {
			ans += v/3 + 1
		} else {
			ans += v / 3
		}

	}

	return ans

}
