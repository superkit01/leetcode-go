package middle

import (
	"slices"
	"sort"
)

func minimumPushes(word string) int {
	wordCnt := make([]int, 26)

	for _, v := range word {
		wordCnt[v-'a']++
	}

	sort.Ints(wordCnt)
	slices.Reverse(wordCnt)

	ans := 0

	for i, v := range wordCnt {
		ans += v * (i/8 + 1)
	}
	return ans
}
