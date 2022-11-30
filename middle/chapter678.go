package middle

func checkValidString(s string) bool {
	cacheStar := make([]int, 0)
	cacheLeft := make([]int, 0)

	for index, v := range s {
		if v == '(' {
			cacheLeft = append(cacheLeft, index)
		}
		if v == '*' {
			cacheStar = append(cacheStar, index)
		}
		if v == ')' {
			if len(cacheLeft) == 0 {
				if len(cacheStar) == 0 {
					return false
				} else {
					cacheStar = cacheStar[0 : len(cacheStar)-1]
				}
			} else {
				cacheLeft = cacheLeft[0 : len(cacheLeft)-1]
			}
		}
	}

	if len(cacheLeft) == 0 {
		return true
	}

	if len(cacheLeft) > len(cacheStar) {
		return false
	}

	for i := len(cacheLeft) - 1; i >= 0; i-- {
		star := cacheStar[len(cacheStar)-1]
		cacheStar = cacheStar[0 : len(cacheStar)-1]
		if star < cacheLeft[i] {
			return false
		}

	}

	return true

}
