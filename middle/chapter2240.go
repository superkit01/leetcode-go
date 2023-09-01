package middle

func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	var sum int64 = 0

	for i := 0; i <= total/cost1; i++ {
		remaining := total - i*cost1
		sum += int64(remaining/cost2 + 1)
	}

	return sum
}
