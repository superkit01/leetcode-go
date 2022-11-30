package easy

func findTheDifference(s string, t string) byte {
	var result rune
	for _, v := range s {
		result ^= v
	}
	for _, v := range t {
		result ^= v
	}
	return byte(result)

}
