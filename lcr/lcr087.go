package lcr

import (
	"strconv"
	"strings"
)

func RestoreIpAddresses(s string) []string {
	if len(s) < 4 {
		return []string{}
	}

	res := [][]string{}

	validate := func(s string) bool {
		if len(s) <= 0 || len(s) >= 4 {
			return false
		}
		if len(s) == 1 {
			return true
		}
		if len(s) == 2 {
			if s[0] == '0' {
				return false
			} else {
				return true
			}
		}
		if s[0] == '0' {
			return false
		}

		i, _ := strconv.Atoi(s)
		return i <= 255
	}

	//0-i闭区间  在j位置划分
	var dfs func(s string, cut int, current []string)
	dfs = func(s string, cut int, current []string) {

		if cut == 0 {
			if validate(s) {
				temp := []string{}
				temp = append(temp, s)
				temp = append(temp, current...)
				res = append(res, temp)
			}
			return
		}

		if len(s) < cut+1 || len(s) > (cut+1)*3 {
			return
		}

		for i := len(s) - 1; i >= len(s)-3 && i >= 0; i-- {
			if validate(s[i:]) {
				temp := []string{}
				temp = append(temp, s[i:])
				temp = append(temp, current...)
				dfs(s[0:i], cut-1, temp)
			}
		}
	}

	dfs(s, 3, []string{})

	ans := []string{}
	for _, v := range res {
		ans = append(ans, strings.Join(v, "."))
	}
	return ans
}
