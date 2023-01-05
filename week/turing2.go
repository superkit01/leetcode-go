package week

func generateParenthesis(n int) []string {
	result := make([]string, 0)

	digui("", n, n, &result)

	return result
}

func digui(str string, left int, right int, result *[]string) {
	if left == 0 && right == 0 {
		*result = append(*result, str)
		return
	}

	if left == right {
		digui(str+"(", left-1, right, result)
	} else if left < right {
		if left > 0 {
			digui(str+"(", left-1, right, result)
		}
		digui(str+")", left, right-1, result)

	}

}
