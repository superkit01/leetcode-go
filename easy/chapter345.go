package main

var Vowels = map[byte]bool{
	'a': true,
	'o': true,
	'e': true,
	'i': true,
	'u': true,
	'A': true,
	'O': true,
	'E': true,
	'I': true,
	'U': true,
}

func reverseVowels(s string) string {
	temp := []byte(s)

	for i, j := 0, len(temp)-1; i < j; {
	inner:
		for i < j {
			if _, ok := Vowels[temp[i]]; !ok {
				i++
			} else {
				break inner
			}
		}

	inner2:
		for i < j {
			if _, ok := Vowels[temp[j]]; !ok {
				j--
			} else {
				break inner2
			}
		}
		temp[i], temp[j] = temp[j], temp[i]
		i++
		j--
	}
	return string(temp)

}
