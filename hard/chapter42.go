package hard

import "math"

func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}

	lmax := make([]int, len(height))
	lmax[0] = height[0]
	for i := 1; i < len(height); i++ {
		lmax[i] = int(math.Max(float64(height[i]), float64(lmax[i-1])))
	}

	rmax := make([]int, len(height))
	rmax[len(height)-1] = height[len(height)-1]

	for i := len(height) - 2; i >= 0; i-- {
		rmax[i] = int(math.Max(float64(height[i]), float64(rmax[i+1])))
	}

	ans := 0

	for i := 1; i < len(height)-1; i++ {
		if int(math.Min(float64(lmax[i-1]), float64(rmax[i+1]))) > height[i] {
			ans += int(math.Min(float64(lmax[i-1]), float64(rmax[i+1]))) - height[i]
		}
	}
	return ans

}
