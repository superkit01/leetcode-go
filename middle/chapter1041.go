package middle

func IsRobotBounded(instructions string) bool {
	x, y := 0, 0
	direct := 'N'

	for _, v := range instructions {

		if direct == 'N' {

			if v == 'G' {
				y++
			}
			if v == 'L' {
				direct = 'W'
			}
			if v == 'R' {
				direct = 'E'
			}
			continue
		}

		if direct == 'W' {

			if v == 'G' {
				x--
			}
			if v == 'L' {
				direct = 'S'
			}
			if v == 'R' {
				direct = 'N'
			}
			continue
		}

		if direct == 'E' {

			if v == 'G' {
				x++
			}
			if v == 'L' {
				direct = 'N'
			}
			if v == 'R' {
				direct = 'S'
			}
			continue
		}

		if direct == 'S' {

			if v == 'G' {
				y--
			}
			if v == 'L' {
				direct = 'E'
			}
			if v == 'R' {
				direct = 'W'
			}
			continue
		}
	}

	if x == 0 && y == 0 || direct != 'N' {
		return true
	} else {
		return false

	}

}
