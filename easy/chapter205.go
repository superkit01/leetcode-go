package main

func isIsomorphic(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	temp := make(map[byte]byte)
	temp2 := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		if _, ok := temp[s[i]]; !ok {
			temp[s[i]] = t[i]

		}
		if _, ok := temp2[t[i]]; !ok {
			temp2[t[i]] = s[i]
		}

		if temp[s[i]] != t[i] || temp2[t[i]] != s[i] {
			return false
		}

	}
	return true

}
