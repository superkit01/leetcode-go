package lcr

func maxProduct(words []string) int {
	wordBit := make([]int, len(words))

	for i, v := range words {
		bit := 0
		for _, w := range v {
			bit |= 1 << (w - 'a')
		}
		wordBit[i] = bit
	}
	max := 0

	for i := 0; i < len(wordBit)-1; i++ {
		for j := i + 1; j < len(wordBit); j++ {
			if wordBit[i]&wordBit[j] == 0 {
				if max < len(words[i])*len(words[j]) {
					max = len(words[i]) * len(words[j])
				}
			}
		}
	}

	return max

}
