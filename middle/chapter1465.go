package middle

import "sort"

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	horizontal := make([]int, 0)
	horizontal = append(horizontal, horizontalCuts[0])
	if len(horizontalCuts) > 1 {
		for i := 1; i < len(horizontalCuts); i++ {
			horizontal = append(horizontal, horizontalCuts[i]-horizontalCuts[i-1])
		}
	}

	horizontal = append(horizontal, h-horizontalCuts[len(horizontalCuts)-1])
	sort.Ints(horizontal)

	vertical := make([]int, 0)
	vertical = append(vertical, verticalCuts[0])
	if len(verticalCuts) > 1 {
		for i := 1; i < len(verticalCuts); i++ {
			vertical = append(vertical, verticalCuts[i]-verticalCuts[i-1])
		}
	}
	vertical = append(vertical, w-verticalCuts[len(verticalCuts)-1])
	sort.Ints(vertical)

	return int((int64(vertical[len(vertical)-1]) * int64(horizontal[len(horizontal)-1])) % 1000000007)

}
