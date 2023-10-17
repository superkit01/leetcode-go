package easy

func sumOfMultiples(n int) int {
	result := 0
	if n >= 3 {
		result += 3 * ((1 + n/3) * (n / 3) / 2)
	}

	if n >= 5 {
		result += 5 * ((1 + n/5) * (n / 5) / 2)
	}

	if n >= 7 {
		result += 7 * ((1 + n/7) * (n / 7) / 2)
	}

	if n >= 15 {
		result += -15 * ((1 + n/15) * (n / 15) / 2)
	}

	if n >= 21 {
		result += -21 * ((1 + n/21) * (n / 21) / 2)
	}

	if n >= 35 {
		result += -35 * ((1 + n/35) * (n / 35) / 2)
	}

	if n >= 105 {
		result += 105 * ((1 + n/105) * (n / 105) / 2)
	}

	return result
}
