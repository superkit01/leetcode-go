package middle

func GarbageCollection(garbage []string, travel []int) int {
	ans := len(garbage[0])
	if len(garbage) == 1 {
		return ans
	}

	posM, posP, posG := 0, 0, 0
	for i := 1; i < len(garbage); i++ {
		for _, v := range garbage[i] {
			if v == 'M' {
				for posM < i {
					ans += travel[posM]
					posM++
				}
			}

			if v == 'P' {
				for posP < i {
					ans += travel[posP]
					posP++
				}
			}

			if v == 'G' {
				for posG < i {
					ans += travel[posG]
					posG++
				}
			}
			ans += 1
		}
	}
	return ans

}
