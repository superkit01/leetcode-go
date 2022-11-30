package easy

func nextGreatestLetter(letters []byte, target byte) byte {

	left := 0
	right := len(letters)

	for left <= right {
		mid := (left + right) / 2

		if letters[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return letters[left]

}
