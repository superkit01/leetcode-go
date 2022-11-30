package easy

//import "fmt"
//
//func main() {
//
//	fmt.Println(climbStairs(2))
//}

var tmp = make(map[int]int)

func climbStairs(n int) int {

	return dp(n)
}

func dp(rest int) int {
	if rest == 2 {
		return 2
	}
	if rest == 1 {
		return 1
	}

	//记忆化搜索  不加执行超时
	if v, ok := tmp[rest]; ok {
		return v
	}

	v1 := dp(rest - 1)
	v2 := dp(rest - 2)

	tmp[rest] = v1 + v2

	return v1 + v2
}
