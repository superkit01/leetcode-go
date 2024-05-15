package lcr

func GenerateParenthesis(n int) []string {
	ans := make([]string, 0)

	current := ""

	dfs085(n, n, current, &ans)

	return ans

}

func dfs085(left int, right int, current string, ans *[]string) {
	if left == 0 && right == 0 {
		*ans = append(*ans, current)
		return
	}

	if left == 0 {
		dfs085(left, right-1, current+")", ans)
		return
	}

	if left == right {
		dfs085(left-1, right, current+"(", ans)
	} else {
		dfs085(left-1, right, current+"(", ans)
		dfs085(left, right-1, current+")", ans)
	}

}
