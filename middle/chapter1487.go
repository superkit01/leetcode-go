package middle

import "fmt"

func getFolderNames(names []string) []string {
	cache := make(map[string]int, 0)

	result := make([]string, 0)
outer:
	for _, v := range names {
		if k, ok := cache[v]; !ok {
			cache[v] = 0
			result = append(result, v)
		} else {
			for i := k + 1; ; i++ {
				tmp := fmt.Sprintf("%v(%v)", v, i)
				if _, ok := cache[tmp]; !ok {
					cache[tmp] = 0
					cache[v] = i
					result = append(result, tmp)
					continue outer
				}
			}
		}
	}

	return result

}
