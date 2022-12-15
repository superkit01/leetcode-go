package easy

import "strconv"

func GetLucky(s string, k int) int {
	target := ""
	for _, v := range s {
		target = target + strconv.Itoa(int(v-'a'+1))
	}

	for i := 0; i < k; i++ {
		if len(target) == 1 {
			k, _ := strconv.Atoi(target)
			return k
		}
		temp := 0
		for _, v := range target {
			temp = temp + int(v-'0')
		}
		target = strconv.Itoa(temp)

	}
	result, _ := strconv.Atoi(target)
	return result
}
