package middle

func wordBreak(s string, wordDict []string) bool {
	cache := map[string]bool{}
	for _, v := range wordDict {
		cache[v] = true
	}

	memo := map[int]bool{}

	var dfs func(index int) bool
	dfs = func(index int) bool {
		if v, ok := memo[index]; ok {
			return v
		}

		if cache[string(s[0:index])] {
			memo[index] = true
			return true
		}

		for _, v := range wordDict {
			if len(s[0:index]) < len(v) {
				continue
			}

			if string(s[index-len(v):index]) == v {
				value := dfs(index - len(v))

				if value {
					memo[index] = true
					return true
				}
			}
		}
		memo[index] = false
		return false
	}

	return dfs(len(s))

}
