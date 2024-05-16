package middle

func NumberOfWeeks(milestones []int) int64 {
	max := int64(0)
	sum := int64(0)
	for _, v := range milestones {
		if max < int64(v) {
			max = int64(v)
		}
		sum += int64(v)
	}

	if max > (sum-max)+int64(1) {
		return (sum-max)*2 + int64(1)
	} else {
		return sum
	}

}
