package main

func minOperations(logs []string) int {
	cache := make([]string, 0)

	for _, v := range logs {
		if v == "./" {
			continue
		} else if v == "../" {
			if len(cache) > 0 {
				cache = cache[0 : len(cache)-1]
			}
		} else {
			cache = append(cache, v)
		}
	}

	return len(cache)

}
