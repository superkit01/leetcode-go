package middle

func minimumRefill(plants []int, capacityA int, capacityB int) int {
	A := 0
	B := len(plants) - 1
	restA := capacityA
	restB := capacityB

	ans := 0

	for A <= B {
		if A == B {
			if restA < restB {
				if restB < plants[B] {
					ans += 1
				}
			} else {
				if restA < plants[A] {
					ans += 1
				}
			}
			return ans
		}

		if restA < plants[A] {
			restA = capacityA
			ans += 1
		}

		restA -= plants[A]
		A++

		if restB < plants[B] {
			restB = capacityB
			ans += 1
		}

		restB -= plants[B]
		B--
	}
	return ans

}
