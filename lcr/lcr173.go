package lcr

func takeAttendance(records []int) int {
	ans := 0
	for i, v := range records {
		ans = ans ^ i ^ v
	}
	ans ^= len(records)
	return ans
}
