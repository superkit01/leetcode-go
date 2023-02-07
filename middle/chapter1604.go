package middle

import (
	"sort"
	"time"
)

func alertNames(keyName []string, keyTime []string) []string {
	group := make(map[string][]string)

	for i := 0; i < len(keyName); i++ {
		if _, ok := group[keyName[i]]; !ok {
			group[keyName[i]] = make([]string, 0)
		}
		group[keyName[i]] = append(group[keyName[i]], keyTime[i])
	}
	result := make([]string, 0)

outer:
	for k, v := range group {
		if len(v) < 3 {
			continue
		}

		sort.Strings(v)

		for i := 2; i < len(v); i++ {
			v1, _ := time.ParseInLocation("15:04", v[i-2], time.Local)
			v2, _ := time.ParseInLocation("15:04", v[i], time.Local)
			duration := v2.Sub(v1)
			if duration.Nanoseconds() <= int64(time.Hour*1) {
				result = append(result, k)
				continue outer
			}
		}

	}
	sort.Strings(result)
	return result

}
