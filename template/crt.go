package template
/**
* Chinese Remainer Theorem
* 三三数之剩二，五五数之剩三，七七数之剩二
*
* [[3,2],[5,3],[7,2]]
*/

func Crt(divRemArr [][]int) int {
	prod := 1
	for _,dr := range divRemArr {
		prod *= dr[0]
	}

	sum := 0
	for _,dr:=range divRemArr {
		n:=prod/dr[0]
		sum += (n * (n%dr[0]) *dr[1]) % prod
	}
	return sum % prod
}
