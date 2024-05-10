package lcr

import "math"

func MinWindow(s string, t string) string {

	if len(t) > len(s) {
		return ""
	}

	cntT := map[rune]int{}
	for _, v := range t {
		cntT[v]++
	}
	cntS := map[rune]int{}

	ansLen := math.MaxInt32
	ans := ""
	i := 0
	j := 0
	for j < len(s) {
		cntS[rune(s[j])]++

		for great(cntS, cntT) && i <= j {
			if ansLen > j+1-i {
				ansLen = j + 1 - i
				ans = s[i : j+1]
			}
			cntS[rune(s[i])]--
			i++

		}
		j++

	}
	return ans

}

func great(cntS, cntT map[rune]int) bool {

	for k, v := range cntT {
		if _, ok := cntS[k]; !ok || cntS[k] < v {
			return false
		}
	}
	return true
}
