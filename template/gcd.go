package template

//最大公约数
// 10  4
// 10 % 4 .. 2
// 4  % 2 .. 0

func gcd(m, n int) int {
	if m%n == 0 {
		return n
	}
	return gcd(n, m%n)
}
