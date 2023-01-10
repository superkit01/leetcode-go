package hard

func longestValidParentheses(s string) int {
	cache := make([]int, 0)
	cache = append(cache, -1)

	result := make([]int, 0)

	for i, v := range s {
		if v == '(' {
			cache = append(cache, i)
		}

		if v == ')' {
			cache = cache[0 : len(cache)-1]
			if len(cache) == 0 {
				cache = append(cache, i)
			} else {
				result = append(result, i-cache[len(cache)-1])
			}
		}

	}

	if len(result) == 0 {
		return 0
	}

	max := result[0]
	for _, v := range result {
		if v > max {
			max = v
		}
	}
	return max
}
