package middle

import "strings"

func AreSentencesSimilar(sentence1 string, sentence2 string) bool {
	s1 := strings.Split(sentence1, " ")
	s2 := strings.Split(sentence2, " ")

	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}

	flag := true
	//头部包含
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			flag = false
			break
		}
	}
	if flag {
		return flag
	}

	//尾部包含
	flag = true
	for i := 0; i < len(s1); i++ {
		if s1[len(s1)-1-i] != s2[len(s2)-1-i] {
			flag = false
			break
		}
	}

	if flag {
		return flag
	}

	i := 0
	j := 0
	for {
		if s1[i] != s2[i] {
			break
		}
		i++
	}

	for {
		if s1[len(s1)-1-j] != s2[len(s2)-1-j] {
			break
		}
		j++
	}
	if i+j < len(s1) {
		flag = false
	} else {
		flag = true
	}

	return flag

}
