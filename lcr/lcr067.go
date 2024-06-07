package lcr

func FindMaximumXOR(nums []int) int {
	trie2 := Constructor2()
	trie2.Insert(nums[0])

	ans := 0
	for i := 1; i < len(nums); i++ {
		maxValue := 0
		current := &trie2
		for j := 31; j >= 0; j-- {
			bit := nums[i] >> j & 0b1
			if bit == 1 {
				if current.left != nil {
					maxValue |= 1 << j
					current = current.left
				} else {
					current = current.right
				}
			}

			if bit == 0 {
				if current.right != nil {
					maxValue |= 1 << j
					current = current.right
				} else {
					current = current.left
				}
			}

		}
		trie2.Insert(nums[i])

		ans = max(ans, maxValue)

	}
	return ans

}

type Trie2 struct {
	left, right *Trie2
}

func Constructor2() Trie2 {
	dummy := Trie2{}
	return dummy

}

func (this *Trie2) Insert(value int) {
	current := this

	for i := 31; i >= 0; i-- {
		bit := (value >> i) & 0b1
		if bit == 0 && current.left == nil {
			current.left = &Trie2{}
		}
		if bit == 1 && current.right == nil {
			current.right = &Trie2{}
		}

		if bit == 0 {
			current = current.left
		} else {
			current = current.right
		}
	}

}

func (this *Trie2) Search(value int) bool {
	current := this
	for i := 31; i >= 0; i-- {
		bit := (value >> i) & 0b1

		if bit == 0 && current.left == nil {
			return false
		}

		if bit == 1 && current.right == nil {
			return false
		}

		if bit == 0 {
			current = current.left
		} else {
			current = current.right
		}
	}
	return true

}
