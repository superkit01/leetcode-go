package lcr

//dp  dp[i]= min(dp[i-1]+1, index- postMap[nums[i]])

func LengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	posMap := make(map[byte]int, 0)

	dp := make([]int, len(s))

	for i := 0; i < len(s); i++ {
		if i == 0 {
			dp[i] = 1
			posMap[s[i]] = 0
		} else {
			if v, ok := posMap[s[i]]; !ok {
				dp[i] = dp[i-1] + 1
			} else {
				dp[i] = min(i-v, dp[i-1]+1)
			}
			posMap[s[i]] = i
		}
	}

	max := 0
	for _, v := range dp {
		if max < v {
			max = v
		}
	}
	return max

}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

//
func LengthOfLongestSubstringII(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	posMap := make(map[byte]int, 0)
	posMap[s[0]]++
	ans := 1

	i := 0
	j := 1
	for j < len(s) {
		for posMap[s[j]] > 0 {
			posMap[s[i]]--
			i++
		}
		posMap[s[j]]++

		if ans < j-i+1 {
			ans = j - i + 1
		}
		j++
	}
	return ans
}
