package middle

import "strconv"

func getHint(secret string, guess string) string {
	x := 0
	y := 0
	cache := make(map[byte][]int, 0)
	skip := make(map[int]bool, 0)
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			x++
			skip[i] = true
		} else {
			cache[secret[i]] = append(cache[secret[i]], i)
		}
	}

	for i := 0; i < len(guess); i++ {
		if _, ok := skip[i]; ok {
			continue
		}

		if _, ok := cache[guess[i]]; ok && len(cache[guess[i]]) > 0 {
			y++
			cache[guess[i]] = cache[guess[i]][1:]
		}

	}
	return strconv.Itoa(x) + "A" + strconv.Itoa(y) + "B"

}
