package middle

import "bytes"

func evaluate(s string, knowledge [][]string) string {
	knowledgeMap := make(map[string]string, 0)
	for _, v := range knowledge {
		knowledgeMap[v[0]] = v[1]
	}

	result := bytes.Buffer{}

	lastLeft := -1
	for i := 0; i < len(s); i++ {

		if s[i] == '(' {
			lastLeft = i
		}

		if lastLeft == -1 {
			result.WriteString(string(s[i]))
		}

		if s[i] == ')' {
			temp := s[lastLeft+1 : i]
			if v, ok := knowledgeMap[temp]; !ok {
				result.WriteString("?")
			} else {
				result.WriteString(v)
			}

			lastLeft = -1

		}
	}

	return result.String()
}
