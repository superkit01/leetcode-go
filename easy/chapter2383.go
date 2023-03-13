package easy

func MinNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	sumEnergy := 0

	for _, v := range energy {
		sumEnergy += v
	}

	increaseEnergy := 0
	if sumEnergy >= initialEnergy {
		increaseEnergy = sumEnergy - initialEnergy + 1

	}

	increaseExperience := 0

	for _, v := range experience {
		tempTncreaseExperience := 0
		if v >= initialExperience {
			tempTncreaseExperience = v - initialExperience + 1
			initialExperience += tempTncreaseExperience
			increaseExperience += tempTncreaseExperience
		}
		initialExperience += v

	}

	return increaseEnergy + increaseExperience
}
