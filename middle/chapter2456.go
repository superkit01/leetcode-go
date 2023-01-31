package middle

func mostPopularCreator(creators []string, ids []string, views []int) [][]string {

	type info struct {
		viewSum int
		maxView int
		maxid   string
	}

	cache := make(map[string]*info, 0)
	maxViewSum := 0

	for i, v := range creators {
		if _, ok := cache[v]; !ok {
			cache[v] = &info{
				viewSum: views[i],
				maxView: views[i],
				maxid:   ids[i],
			}
		} else {
			cache[v].viewSum = cache[v].viewSum + views[i]
			if cache[v].maxView < views[i] {
				cache[v].maxView = views[i]
				cache[v].maxid = ids[i]
			} else if cache[v].maxView == views[i] {
				if ids[i] < cache[v].maxid {
					cache[v].maxid = ids[i]
				}
			}
		}

		if cache[v].viewSum > maxViewSum {
			maxViewSum = cache[v].viewSum
		}
	}

	result := make([][]string, 0)

	for k, v := range cache {
		if v.viewSum == maxViewSum {
			temp := make([]string, 0)
			temp = append(temp, k)
			temp = append(temp, v.maxid)
			result = append(result, temp)
		}
	}

	return result

}
