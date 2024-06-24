package lcp

func temperatureTrend(temperatureA []int, temperatureB []int) int {
	ans := make([]int, len(temperatureA))
	for i := 1; i < len(temperatureA); i++ {
		statusA := 0
		if temperatureA[i]-temperatureA[i-1] > 0 {
			statusA = 1
		} else if temperatureA[i]-temperatureA[i-1] < 0 {
			statusA = -1
		}

		statusB := 0
		if temperatureB[i]-temperatureB[i-1] > 0 {
			statusB = 1
		} else if temperatureB[i]-temperatureB[i-1] < 0 {
			statusB = -1
		}

		if statusA == statusB {
			ans[i] = ans[i-1] + 1
		}

	}

	max := 0
	for _, v := range ans {
		if max < v {
			max = v
		}

	}
	return max

}
