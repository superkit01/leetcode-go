package middle

import (
	"math"
	"strings"
)

func AlphabetBoardPath(target string) string {
	cache := make(map[rune][]int, 0)
	cache['a'] = []int{0, 0}
	cache['b'] = []int{0, 1}
	cache['c'] = []int{0, 2}
	cache['d'] = []int{0, 3}
	cache['e'] = []int{0, 4}

	cache['f'] = []int{1, 0}
	cache['g'] = []int{1, 1}
	cache['h'] = []int{1, 2}
	cache['i'] = []int{1, 3}
	cache['j'] = []int{1, 4}

	cache['k'] = []int{2, 0}
	cache['l'] = []int{2, 1}
	cache['m'] = []int{2, 2}
	cache['n'] = []int{2, 3}
	cache['o'] = []int{2, 4}

	cache['p'] = []int{3, 0}
	cache['q'] = []int{3, 1}
	cache['r'] = []int{3, 2}
	cache['s'] = []int{3, 3}
	cache['t'] = []int{3, 4}

	cache['u'] = []int{4, 0}
	cache['v'] = []int{4, 1}
	cache['w'] = []int{4, 2}
	cache['x'] = []int{4, 3}
	cache['y'] = []int{4, 4}

	cache['z'] = []int{5, 0}

	pre := []int{0, 0}
	result := ""
	for _, v := range target {
		result = result + calcute(pre, cache[v])
		pre = cache[v]
	}

	return result

}

func calcute(from []int, to []int) string {
	result := ""
	xAxis := to[1] - from[1]
	yAxis := to[0] - from[0]
	if xAxis == 0 && yAxis == 0 {
		result = result + "!"
		return result
	}
	if yAxis < 0 {
		result = result + strings.Repeat("U", int(math.Abs(float64(yAxis))))
	}
	if xAxis < 0 {
		result = result + strings.Repeat("L", int(math.Abs(float64(xAxis))))
	}
	if yAxis > 0 {
		result = result + strings.Repeat("D", int(math.Abs(float64(yAxis))))
	}

	if xAxis > 0 {
		result = result + strings.Repeat("R", int(math.Abs(float64(xAxis))))
	}

	result = result + "!"
	return result
}
