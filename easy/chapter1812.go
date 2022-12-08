package easy

import "strconv"

func squareIsWhite(coordinates string) bool {
	cache := map[byte]int{
		'a': 1,
		'b': 2,
		'c': 3,
		'd': 4,
		'e': 5,
		'f': 6,
		'g': 7,
		'h': 8,
	}

	v1, _ := cache[coordinates[0]]
	v2, _ := strconv.Atoi(string(coordinates[1]))

	if (v1+v2)%2 == 0 {
		return false
	} else {
		return true
	}

}
