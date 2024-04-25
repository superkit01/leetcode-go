package middle

import "sort"

func eliminateMaximum(dist []int, speed []int) int {
	time := make([]float64, len(dist))
	for i := 0; i < len(time); i++ {
		time[i] = float64(dist[i]) / float64(speed[i])
	}

	sort.Float64s(time)
	i := 0
	for i < len(time)-1 {
		i++
		if time[i]-float64(i) <= 0 {
			return i
		}

	}
	return i + 1
}
