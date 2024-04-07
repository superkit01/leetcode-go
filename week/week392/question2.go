package week392

import "math"

func GetSmallestString(s string, k int) string {

	ans := ""

	for _, v := range s {
		if k == 0 {
			ans += string(v)
			continue
		}
		dis := int(math.Min(math.Abs(float64(v-'a')), math.Abs(float64('a'+26-v))))
		if dis <= k {
			k = k - dis
			ans += "a"
		} else {
			ans += string(rune(v - rune(k)))
			k = 0
		}
	}
	return ans

}
