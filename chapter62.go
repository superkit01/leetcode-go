package main

func uniquePaths(m int, n int) int {

	cache := make(map[int]map[int]int, 0)
	return dp2(m-1, n-1, cache)

}

func dp2(restM int, restN int, cache map[int]map[int]int) int {

	if restM == 0 || restN == 0 {
		return 1
	}

	if k, ok := cache[restM]; ok {
		if v, ok := k[restN]; ok {
			return v
		}
	}

	result := dp2(restM-1, restN, cache) + dp2(restM, restN-1, cache)

	if _, ok := cache[restM]; !ok {
		cache[restM] = make(map[int]int)
	}
	temp := cache[restM]
	if _, ok := temp[restN]; !ok {
		temp[restN] = 0
	}
	temp[restN] = result

	return result
}

//func main() {
//
//	fmt.Println(uniquePaths(51, 9))
//}
