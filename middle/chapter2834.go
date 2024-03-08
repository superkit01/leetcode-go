package middle

func minimumPossibleSum(n int, target int) int {
	var sum int64 = 0
	if n <= target/2 {
		sum = int64(n) * int64((1 + n)) / 2
		return int(sum % 1000000007)
	} else {
		sum += int64(target/2) * int64((1 + target/2)) / 2
		sum += (int64(target) + int64(target+(n-target/2)-1)) * int64(n-target/2) / 2
		return int(sum % 1000000007)
	}

}
