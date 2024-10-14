package lcr

// 整数数组 sockets 记录了一个袜子礼盒的颜色分布情况，其中 sockets[i] 表示该袜子的颜色编号。
// 礼盒中除了一款撞色搭配的袜子，每种颜色的袜子均有两只。
// 请设计一个程序，在时间复杂度 O(n)，空间复杂度O(1) 内找到这双撞色搭配袜子的两个颜色编号。

func sockCollocation(sockets []int) []int {

	xor := 0
	for _, v := range sockets {
		xor ^= v
	}

	i := 1

	for xor&i != i {
		i <<= 1
	}

	ans1 := 0
	ans2 := 0
	for _, v := range sockets {
		if v&i == i {
			ans1 ^= v
		} else {
			ans2 ^= v
		}
	}
	return []int{ans1, ans2}
}
