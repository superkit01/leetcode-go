package middle

func minimumSwap(s1 string, s2 string) int {
	countXY := 0
	countYX := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			continue
		}
		if s1[i] == 'x' {
			countXY++
		} else {
			countYX++
		}
	}
	if (countXY+countYX)%2 != 0 {
		return -1
	}
	return countXY/2 + countYX/2 + (countXY%2)*2

}
