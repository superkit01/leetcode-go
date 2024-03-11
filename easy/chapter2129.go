package easy

import (
	"strings"
)

func capitalizeTitle(title string) string {
	strs := strings.Split(title, " ")
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) <= 2 {
			strs[i] = strings.ToLower(strs[i])
			continue
		}
		strs[i] = strings.ToUpper(strs[i][:1]) + strings.ToLower(strs[i][1:])
	}

	return strings.Join(strs, " ")

}
