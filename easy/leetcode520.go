package easy

func detectCapitalUse(word string) bool {
	if len(word) == 1 {
		return true
	}

	if word[0] >= 'a' && word[0] <= 'z' {
		for i := 0; i < len(word); i++ {
			if word[i] < 'a' || word[i] > 'z' {
				return false
			}
		}
		return true
	} else {
		i := 1
		for i < len(word)-1 {
			if (word[i] >= 'a' && word[i] <= 'z') && (word[i+1] >= 'a' && word[i+1] <= 'z') {
				i++
				continue
			}
			if (word[i] < 'a' || word[i] > 'z') && (word[i+1] < 'a' || word[i+1] > 'z') {
				i++
				continue
			}
			i++
			return false
		}

		return true
	}

}
