package middle

import "bytes"

func Evaluate(s string, knowledge [][]string) string {
	kw := make(map[string]string, 0)
	for _, v := range knowledge {
		kw[v[0]] = v[1]
	}

	i := -1
	ans := bytes.NewBufferString("")

	for j := 0; j < len(s); j++ {
		if s[j] == '(' {
			i = j + 1
			continue
		}
		if s[j] == ')' {
			key := s[i:j]
			if v, ok := kw[key]; ok {
				ans.WriteString(v)
			} else {
				ans.WriteRune('?')
			}

			i = -1
			continue
		}
		if i == -1 {
			ans.WriteByte(s[j])
		}

	}

	return ans.String()

}
