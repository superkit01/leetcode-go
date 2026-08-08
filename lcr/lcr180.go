package lcr

func FileCombination(target int) [][]int {
	i, j := 1, 2

	if target < 3 {
		return [][]int{{1, 2}}
	}
	ans := make([][]int, 0)

	sum := 3

	for i <= target/2 && i < j {

		if sum < target {
			j++
			sum += j
		} else if sum > target {
			sum -= i
			i++
		} else {
			res := make([]int, j-i+1)
			for index := range res {
				res[index] = i + index
			}
			ans = append(ans, res)
			sum -= i
			i++

		}
	}
	return ans

}
