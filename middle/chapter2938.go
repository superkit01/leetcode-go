package middle

func minimumSteps(s string) int64 {
	ans := int64(0)

	i := 0
	j := len(s) - 1

	for i < j {
		for s[i] == '0' && i < j {
			i++
		}

		for s[j] == '1' && i < j {
			j--
		}

		if i >= j {
			break
		}
		ans += int64(j - i)
		i++
		j--
	}
	return ans
}
