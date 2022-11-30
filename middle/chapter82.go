package middle

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cache := make(map[int]int)
	cur := head
	for cur != nil {
		if _, ok := cache[cur.Val]; !ok {
			cache[cur.Val] = 0
		}
		cache[cur.Val]++
		cur = cur.Next
	}

	var result = &ListNode{Val: -1, Next: nil}
	temp := result
	cur = head
	for cur != nil {
		v, _ := cache[cur.Val]
		if v == 1 {
			temp.Next = &ListNode{cur.Val, nil}
			temp = temp.Next
		}
		cur = cur.Next
	}

	return result.Next

}

// func main() {
// 	temp := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, nil}}}}}}}
// 	deleteDuplicates3(temp)
// }
