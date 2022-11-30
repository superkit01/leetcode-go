package easy

func reformat(s string) string {
	charactorQueue := make([]rune, 0)
	digistQueue := make([]rune, 0)

	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			charactorQueue = append(charactorQueue, v)
		} else {
			digistQueue = append(digistQueue, v)
		}
	}

	if len(charactorQueue) >= len(digistQueue)+2 ||
		len(digistQueue) >= len(charactorQueue)+2 {
		return ""
	}

	result := ""
	charStart := false
	if len(charactorQueue) > len(digistQueue) {
		charStart = true
	}

	charI := 0
	digI := 0
	for {
		if digI >= len(digistQueue) && charI >= len(charactorQueue) {
			break
		}
		if charStart {
			result = result + string(charactorQueue[charI])
			charI += 1
			charStart = !charStart
		} else {
			result = result + string(digistQueue[digI])
			digI += 1
			charStart = !charStart
		}
	}
	return result

}
