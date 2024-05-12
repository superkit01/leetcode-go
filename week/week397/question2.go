package week397

import "math"

func maximumEnergy(energy []int, k int) int {
	ans := math.MinInt32
	cache := make(map[int]int, 0)
	for i := 0; i < len(energy); i++ {
		value := dfs(energy, i, k, &cache)
		if ans < value {
			ans = value
		}
	}
	return ans

}

func dfs(energy []int, i int, k int, cache *map[int]int) int {
	if v, ok := (*cache)[i]; ok {
		return v
	}

	if i >= len(energy) {
		(*cache)[i] = 0
		return 0
	}
	temp := dfs(energy, i+k, k, cache) + energy[i]
	(*cache)[i] = temp
	return temp
}

func maximumEnergyII(energy []int, k int) int {
	for j := len(energy) - 1; j >= len(energy)-k; j-- {
		for i := j; i >= 0; i -= k {
			if i != j {
				energy[i] = energy[i] + energy[i+k]
			}
		}
	}

	max := math.MinInt32
	for _, v := range energy {
		if max < v {
			max = v
		}
	}

	return max

}
