package easy

func hasCycle(head *ListNode) bool {
	cache := make(map[*ListNode]int)
	for head != nil {
		if _, ok := cache[head]; ok {
			return true
		} else {
			cache[head] = 1
			head = head.Next
		}
	}
	return false

}
