package lcr

func findTargetIn2DPlants(plants [][]int, target int) bool {
	//从右上角观察，矩阵可以转换成BST

	if len(plants) == 0 || len(plants[0]) == 0 {
		return false
	}

	m := 0
	n := len(plants[0]) - 1

	for m < len(plants) && n >= 0 {

		if plants[m][n] < target {
			m += 1
			continue
		}
		if plants[m][n] > target {
			n -= 1
			continue
		}

		if plants[m][n] == target {
			return true
		}

	}
	return false

}
