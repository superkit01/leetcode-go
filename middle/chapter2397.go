package middle

func MaximumRows(matrix [][]int, numSelect int) int {

	rows := make([]int, len(matrix))

	for m := 0; m < len(matrix); m++ {
		for n := 0; n < len(matrix[0]); n++ {
			if matrix[m][n] == 0 {
				continue
			}
			rows[m] |= (1 << n)
		}
	}

	return dp(rows, 0, numSelect, len(matrix[0]))

}

func dp(rows []int, current int, remaining int, max int) int {

	if remaining == 0 {
		return count0(rows)
	} else {
		if current < max {

			newRows := make([]int, 0)
			newRows = append(newRows, rows...)
			for i := 0; i < len(newRows); i++ {
				newRows[i] &= ^(1 << current)
			}

			notUse := dp(rows, current+1, remaining, max)

			use := dp(newRows, current+1, remaining-1, max)

			if notUse > use {
				return notUse
			} else {
				return use
			}

		} else {
			return -1
		}

	}

}

func count0(rows []int) int {
	count := 0

	for _, i := range rows {
		if i == 0 {
			count++
		}
	}
	return count

}
