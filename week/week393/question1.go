package week393

func FindLatestTime(s string) string {
	ans := ""
	if s[0] == '?' {
		if s[1] == '?' {
			ans += "11"
		} else if s[1] > '1' {
			ans += "0"
			ans += string(s[1])
		} else {
			ans += "1"
			ans += string(s[1])
		}
	} else {
		if s[0] == '0' {
			if s[1] == '?' {
				ans += "09"
			} else {
				ans += "0"
				ans += string(s[1])
			}
		} else {
			if s[1] == '?' {
				ans += "11"
			} else {
				ans += "1"
				ans += string(s[1])
			}
		}
	}

	ans += string(':')

	if s[3] == '?' {
		if s[4] == '?' {
			ans += "59"
		} else {
			ans += "5"
			ans += string(s[4])
		}
	} else {
		ans += string(s[3])
		if s[4] == '?' {
			ans += "9"
		} else {
			ans += string(s[4])
		}
	}
	return ans

}
