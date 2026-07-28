package lcr

func trainingPlanI(actions []int) []int {
	i, j := 0, len(actions)-1

	for i < j {
		if actions[i]%2 == 1 {
			i++
			continue
		}

		if actions[j]%2 == 0 {
			j--
			continue
		}

		actions[i], actions[j] = actions[j], actions[i]

	}
	return actions

}
