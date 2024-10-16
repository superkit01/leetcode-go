package lcr

func statisticalResult(arrayA []int) []int {
	mulpy := 1
	count0 := 0
	for _, v := range arrayA {

		if v == 0 {
			count0++
			continue
		}
		mulpy *= v
	}

	for i := range arrayA {

		if count0 > 1 {
			arrayA[i] = 0
			continue
		}

		if count0 == 1 {
			if arrayA[i] != 0 {
				arrayA[i] = 0
			} else {
				arrayA[i] = mulpy
			}
			continue
		}
		arrayA[i] = mulpy / arrayA[i]
	}
	return arrayA

}
