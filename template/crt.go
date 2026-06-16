package template

/**
* Chinese Remainder Theorem 中国剩余定理
* 三三数之剩二，五五数之剩三，七七数之剩二
* 求解同余方程组：x ≡ a1 (mod m1), x ≡ a2 (mod m2), ... (mi 两两互质)
*
* 输入 [[m1,a1],[m2,a2],...]，例如 [[3,2],[5,3],[7,2]] => 23
*
* 公式：M = ∏mi,  Mi = M/mi,  Mi' = Mi 关于 mi 的模逆元
*      x = (∑ ai * Mi * Mi') mod M
*/

// extGcd 扩展欧几里得算法
// 返回 gcd(a,b)，以及满足 a*x + b*y = gcd(a,b) 的 (x,y)
func extGcd(a, b int) (int, int, int) {
	if b == 0 {
		return a, 1, 0
	}
	g, x, y := extGcd(b, a%b)
	return g, y, x-(a/b)*y
}

// modInverse 求 a 关于 m 的模逆元（要求 a 与 m 互质）
// 即满足 a*inv ≡ 1 (mod m) 的最小非负整数
func modInverse(a, m int) int {
	_, x, _ := extGcd(a, m)
	return ((x % m) + m) % m
}

func Crt(divRemArr [][]int) int {
	prod := 1
	for _, dr := range divRemArr {
		prod *= dr[0]
	}

	sum := 0
	for _, dr := range divRemArr {
		m, a := dr[0], dr[1]
		n := prod / m           // Mi = M / mi
		inv := modInverse(n, m) // Mi 关于 mi 的模逆元
		sum += (a * n * inv) % prod
	}
	return sum % prod
}
