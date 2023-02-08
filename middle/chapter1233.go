package middle

import (
	"sort"
	"strings"
)

func removeSubfolders(folder []string) []string {
	sort.Strings(folder)
	result := make([]string, 0)
	result = append(result, folder[0])

	if len(folder) == 1 {
		return result
	}
	j := 0
	for i := 1; i < len(folder); i++ {
		if strings.HasPrefix(folder[i], folder[j]+"/") {
			continue
		}
		j = i
		result = append(result, folder[i])
	}

	return result

}
