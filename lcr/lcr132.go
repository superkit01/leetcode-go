package lcr

func cuttingBambooII(bamboo_len int) int {
	if bamboo_len < 4 {
		return 1 * (bamboo_len - 1)
	}
	if bamboo_len == 4 {
		return 4
	}

	if bamboo_len%3 == 1 {
		return 4 * quickX(3, (bamboo_len-4)/3)
	} else if bamboo_len%3 == 2 {
		return 2 * quickX(3, (bamboo_len-2)/3)
	} else {
		return quickX(3, bamboo_len/3)
	}

}

func quickX(x, n int) int {
	if n == 0 {
		return 1
	}

	if n%2 == 1 {
		return quickX(x, n-1) * x % 1000000007
	}

	return quickX(x, n/2) * quickX(x, n/2) % 1000000007

}
