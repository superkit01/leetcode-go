package middle

// dp [i][j]
//  s[i:j] 是否是回文串

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}

	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
outer:
	for i := 1; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			if i == j {
				dp[i][j] = true
				continue outer
			}
			if i-j == 1 {
				if s[i] == s[j] {
					dp[i][j] = true
				} else {
					dp[i][j] = false
				}
				continue outer
			}

			if dp[i-1][j+1] && s[i] == s[j] {
				dp[i][j] = true
			} else {
				dp[i][j] = false
			}
		}
	}

	ans := 0
	max := s[0:0]
	for i := len(s) - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			if dp[i][j] {
				if ans < i-j+1 {
					ans = i - j + 1
					max = s[j : i+1]
				}
			}
		}
	}
	return max

}
