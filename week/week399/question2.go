package week399

import (
	"bytes"
	"strconv"
)

func compressedString(word string) string {
	cnt := make([][2]int, 0)
	cnt = append(cnt, [2]int{int(word[0]), 1})

	if len(word) > 1 {
		for i := 1; i < len(word); i++ {
			pop := cnt[len(cnt)-1]
			if int(word[i]) == pop[0] && pop[1] < 9 {
				cnt[len(cnt)-1][1]++
			} else {
				cnt = append(cnt, [2]int{int(word[i]), 1})
			}

		}
	}

	var buffer bytes.Buffer
	for _, v := range cnt {
		buffer.WriteString(strconv.Itoa(v[1]))
		buffer.WriteString(string(rune(v[0])))

	}
	return buffer.String()
}
