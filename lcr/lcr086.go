package lcr

func Partition(s string) [][]string {
	ans := make([][]string, 0)
	current := make([]string, 0)

	cache := make(map[string]bool)

	dfs086(s, len(s)-1, &ans, current, &cache)
	return ans
}

func dfs086(s string, index int, ans *[][]string, current []string, cache *map[string]bool) {

	if index < 0 {
		*ans = append(*ans, current)
		return
	}

	for j := index; j >= 0; j-- {
		if isPalindrome(s[j:index+1], cache) {
			temp := []string{}
			temp = append(temp, s[j:index+1])
			temp = append(temp, current...)
			dfs086(s, j-1, ans, temp, cache)
		}
	}

}

func isPalindrome(s string, cache *map[string]bool) bool {
	if v, ok := (*cache)[s]; ok {
		return v
	}
	if len(s) == 0 {
		(*cache)[s] = false
		return false
	}
	if len(s) == 1 {
		(*cache)[s] = true
		return true
	}

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			(*cache)[s] = false
			return false
		}
	}
	(*cache)[s] = true
	return true
}
