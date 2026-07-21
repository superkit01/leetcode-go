package lcr

/*
*

给定两个 非空链表 l1和 l2 来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。

可以假设除了数字 0 之外，这两个数字都不会以零开头。

  - Definition for singly-linked list.
  - type ListNode struct {
  - Val int
  - Next *ListNode
  - }
*/

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverse(l1)
	l2 = reverse(l2)

	ans := &ListNode{}
	curr := ans

	curr1 := l1
	curr2 := l2
	carry := 0

	for curr1 != nil || curr2 != nil || carry != 0 {

		if curr1 != nil {
			carry += curr1.Val
		}

		if curr2 != nil {
			carry += curr2.Val
		}

		curr.Next = &ListNode{
			Val:  carry % 10,
			Next: nil,
		}
		curr = curr.Next

		carry /= 10
		if curr1 != nil {
			curr1 = curr1.Next
		}
		if curr2 != nil {
			curr2 = curr2.Next
		}
	}

	return reverse(ans.Next)
}

func reverse(node *ListNode) *ListNode {

	if node == nil || node.Next == nil {
		return node
	}

	head := reverse(node.Next)
	tail := node.Next
	tail.Next = node
	node.Next = nil

	return head

}
