package easy

func numberOfLines(widths []int, s string) []int {
	total := 0
	line := 1
	for _, c := range s {
		index := c - 'a'
		if total+widths[index] > 100 {
			total = widths[index]
			line++
		} else {
			total += widths[index]
		}
	}
	return []int{line, total}
}
