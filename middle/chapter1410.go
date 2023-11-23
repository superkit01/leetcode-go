package middle

func entityParser(text string) string {
	result := make([]rune, 0)
	for i := 0; i < len(text); {
		if text[i] == '&' {
			if i < len(text)-3 {
				if text[i+1] == 'g' && text[i+2] == 't' && text[i+3] == ';' {
					i += 4
					result = append(result, '>')
					continue
				}
				if text[i+1] == 'l' && text[i+2] == 't' && text[i+3] == ';' {
					i += 4
					result = append(result, '<')
					continue
				}
			}
			if i < len(text)-4 {
				if text[i+1] == 'a' && text[i+2] == 'm' && text[i+3] == 'p' && text[i+4] == ';' {
					i += 5
					result = append(result, '&')
					continue
				}
			}
			if i < len(text)-5 {
				if text[i+1] == 'a' && text[i+2] == 'p' && text[i+3] == 'o' && text[i+4] == 's' && text[i+5] == ';' {
					i += 6
					result = append(result, '\'')
					continue
				}
				if text[i+1] == 'q' && text[i+2] == 'u' && text[i+3] == 'o' && text[i+4] == 't' && text[i+5] == ';' {
					i += 6
					result = append(result, '"')
					continue
				}
			}

			if i < len(text)-6 {
				if text[i+1] == 'f' && text[i+2] == 'r' && text[i+3] == 'a' && text[i+4] == 's' && text[i+5] == 'l' && text[i+6] == ';' {
					i += 7
					result = append(result, '/')
					continue
				}
			}

			result = append(result, rune(text[i]))
			i += 1
		} else {
			result = append(result, rune(text[i]))
			i += 1
		}
	}
	return string(result)

}
