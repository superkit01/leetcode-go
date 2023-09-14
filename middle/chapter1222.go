package middle

func QueensAttacktheKing(queens [][]int, king []int) [][]int {
	x_p := make([]int, 0)
	x_n := make([]int, 0)
	y_p := make([]int, 0)
	y_n := make([]int, 0)
	x_p_y_p := make([]int, 0)
	x_p_y_n := make([]int, 0)
	x_n_y_p := make([]int, 0)
	x_n_y_n := make([]int, 0)

	for _, v := range queens {
		if v[0] == king[0] {
			if v[1] > king[1] {
				if len(y_n) == 0 || v[1] < y_n[1] {
					y_n = v
				}
			} else {
				if len(y_p) == 0 || v[1] > y_p[1] {
					y_p = v
				}
			}
		}

		if v[1] == king[1] {
			if v[0] > king[0] {
				if len(x_p) == 0 || v[0] < x_p[0] {
					x_p = v
				}
			} else {
				if len(x_n) == 0 || v[0] > x_n[0] {
					x_n = v
				}
			}
		}

		if v[0]-v[1] == king[0]-king[1] {
			if v[0] > king[0] {
				if len(x_p_y_n) == 0 || v[0] < x_p_y_n[0] {
					x_p_y_n = v
				}
			} else {
				if len(x_n_y_p) == 0 || v[0] > x_n_y_p[0] {
					x_n_y_p = v
				}
			}
		}

		if v[0]+v[1] == king[0]+king[1] {
			if v[0] > king[0] {
				if len(x_p_y_p) == 0 || v[0] < x_p_y_p[0] {
					x_p_y_p = v
				}
			} else {
				if len(x_n_y_n) == 0 || v[0] > x_n_y_n[0] {
					x_n_y_n = v
				}
			}
		}

	}

	result := make([][]int, 0)
	if len(x_p) > 0 {
		result = append(result, x_p)
	}
	if len(x_n) > 0 {
		result = append(result, x_n)
	}
	if len(y_p) > 0 {
		result = append(result, y_p)
	}
	if len(y_n) > 0 {
		result = append(result, y_n)
	}
	if len(x_p_y_p) > 0 {
		result = append(result, x_p_y_p)
	}
	if len(x_p_y_n) > 0 {
		result = append(result, x_p_y_n)
	}
	if len(x_n_y_p) > 0 {
		result = append(result, x_n_y_p)
	}
	if len(x_n_y_n) > 0 {
		result = append(result, x_n_y_n)
	}

	return result

}
