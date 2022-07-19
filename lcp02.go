package main

func fraction(cont []int) []int {
	numerator := 1
	denominator := cont[len(cont)-1]

	for i := len(cont) - 1; i > 0; i-- {
		temp := numerator
		numerator = denominator
		denominator = denominator*cont[i-1] + temp
	}
	maxV := gcd(numerator, denominator)
	return []int{denominator / maxV, numerator / maxV}
}

func gcd(numerator, denominator int) int {
	for denominator != 0 {
		temp := denominator
		denominator = numerator % denominator
		numerator = temp
	}
	return numerator
}
