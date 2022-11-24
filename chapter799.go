package main

func champagneTower(poured int, query_row int, query_glass int) float64 {

	line := []float64{float64(poured)}
	for row := 0; row < query_row; row++ {
		newLine := make([]float64, len(line)+1)
		for i := 0; i < len(newLine); i++ {
			if i == 0 {
				if line[i] < 1 {
					newLine[i] = 0
				} else {
					newLine[i] = (line[i] - 1) / 2
				}
			} else if i == len(newLine)-1 {
				if line[i-1] < 1 {
					newLine[i] = 0
				} else {
					newLine[i] = (line[i-1] - 1) / 2
				}
			} else {
				var left float64 = 0
				if line[i-1] > 1 {
					left = (line[i-1] - 1) / 2
				}
				var right float64 = 0
				if line[i] > 1 {
					right = (line[i] - 1) / 2
				}
				newLine[i] = left + right
			}
		}
		line = newLine
	}
	if line[query_glass] < 0 {
		return 0
	} else if line[query_glass] >= 1 {
		return 1
	} else {
		return line[query_glass]
	}
}
