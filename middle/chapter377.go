package middle

//dfs + 记忆化搜索
func combinationSum4(nums []int, target int) int {
	cache := make(map[int]int, 0)
	return dfs377(nums, target, &cache)
}

func dfs377(nums []int, remaining int, cache *map[int]int) int {
	if remaining < 0 {
		return 0
	}
	if remaining == 0 {
		return 1
	}
	//记忆化搜索
	if v, ok := (*cache)[remaining]; ok {
		return v
	}

	ans := 0
	for i := 0; i < len(nums); i++ {
		ans += dfs377(nums, remaining-nums[i], cache)
	}
	(*cache)[remaining] = ans
	return ans

}
