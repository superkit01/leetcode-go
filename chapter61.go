package main

func rotateRight(head *ListNode, k int) *ListNode {

	if head == nil || k == 0 || head.Next == nil {
		return head
	}

	iter := head
	length := 1
	for iter.Next != nil {
		length += 1
		iter = iter.Next
	}

	remainder := k % length

	if remainder == 0 {
		return head
	}

	iter.Next = head

	for i := 0; i < length-remainder; i++ {
		iter = iter.Next
	}
	result := iter.Next
	iter.Next = nil
	return result

}

// func main() {
// 	head := &ListNode{
// 		Val: 1,
// 		Next: &ListNode{
// 			Val: 2,
// 			Next: &ListNode{
// 				Val: 3,
// 				Next: &ListNode{
// 					Val: 4,
// 					Next: &ListNode{
// 						Val:  5,
// 						Next: nil,
// 					},
// 				},
// 			},
// 		},
// 	}
// 	rotateRight(head, 5)
// }
