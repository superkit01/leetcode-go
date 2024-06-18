package middle

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func discountPrices(sentence string, discount int) string {
	vacab := strings.Split(sentence, " ")
	ans := []string{}
outer:
	for _, v := range vacab {
		if v[0] != '$' || len(v) == 1 {
			ans = append(ans, v)
			continue
		}

		for i := 1; i < len(v); i++ {
			if !unicode.IsDigit(rune(v[i])) {
				ans = append(ans, v)
				continue outer
			}
		}

		i, _ := strconv.Atoi(v[1:])
		ans = append(ans, fmt.Sprintf("$%.2f", float64(i)*float64(100-discount)/float64(100)))
	}

	return strings.Join(ans, " ")
}
