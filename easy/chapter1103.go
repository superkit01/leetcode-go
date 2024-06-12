package easy

func DistributeCandies(candies int, num_people int) []int {
	//  (1+n+ 2*i*n)*n/2
	turns := binarySeach(candies, num_people)

	ans := make([]int, num_people)

	if turns > 0 {
		for i := 0; i < num_people; i++ {
			ans[i] = (i + 1 + num_people*(turns-1) + i + 1) * turns / 2
		}
	}

	start := num_people * (1 + num_people) / 2
	end := num_people * (1 + num_people + 2*(turns-1)*num_people) / 2
	remain := candies - (start+end)*turns/2
	for i := 0; i < num_people && remain > 0; i++ {
		current := min(turns*num_people+i+1, remain)
		remain -= current
		ans[i] += current
	}
	return ans

}

func binarySeach(candies int, num_people int) int {
	l := 0
	r := candies/num_people + 1
	start := num_people * (1 + num_people) / 2
	for l < r {
		mid := (l + r) / 2
		end := num_people * (1 + num_people + 2*mid*num_people) / 2
		sum := (start + end) * (mid + 1) / 2
		if sum <= candies {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
