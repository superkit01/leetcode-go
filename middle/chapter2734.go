package middle

import "bytes"

func smallestString(s string) string {
	var ans bytes.Buffer
	index := -1
	for i, v := range s {
		if v == 'a' {
			ans.WriteRune(v)
			continue
		}

		if index == i-1 || index == -1 {
			ans.WriteRune(v - 1)
			index = i
			continue
		}
		ans.WriteRune(v)
	}

	if index != -1 {
		return ans.String()

	}
	s = ans.String()
	return s[0:len(s)-1] + string('z')

}
