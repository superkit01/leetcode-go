package middle

//    -1 | 0 1 2 3
//		 | 2 2 3 3	~	 5
//
func wateringPlants(plants []int, capacity int) int {
	ans := 0

	r := 0
	sum := plants[r]
	for {
		for sum <= capacity {
			if r == len(plants)-1 {
				ans += len(plants)
				return ans
			}

			r++
			sum += plants[r]
		}

		ans += 2 * r
		sum = plants[r]
	}

}
