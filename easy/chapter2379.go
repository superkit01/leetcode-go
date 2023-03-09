package easy

func MinimumRecolors(blocks string, k int) int {
	countW := make([]int, 0)

	curW := 0
	for j := 0; j < k; j++ {
		if blocks[j] == 'W' {
			curW++
		}
	}
	countW = append(countW, curW)

	if len(blocks) == k {
		return curW
	}

	for i := 1; i <= len(blocks)-k; i++ {
		if blocks[i-1] == 'W' {
			curW--
		}
		if blocks[i+k-1] == 'W' {
			curW++
		}
		countW = append(countW, curW)
	}

	minW := k
	for _, v := range countW {
		if minW > v {
			minW = v
		}
	}
	return minW

}
