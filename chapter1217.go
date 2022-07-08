package main

func minCostToMoveChips(position []int) int {
	var oddSum int = 0
	var evenSum int = 0
	for _, e := range position {
		if e%2 == 0 {
			oddSum += 1
		} else {
			evenSum += 1
		}
	}

	if oddSum > evenSum {
		return evenSum
	} else {
		return oddSum
	}

}
