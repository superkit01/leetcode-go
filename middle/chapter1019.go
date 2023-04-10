package middle

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nextLargerNodes(head *ListNode) []int {

	temp := make([]int, 0)
	for head != nil {
		temp = append(temp, head.Val)
		head = head.Next
	}

	stack := make([]int, 0)

	result := make([]int, len(temp))

outer:
	for i := len(temp) - 1; i >= 0; i-- {

		for {
			if len(stack) == 0 {
				result[i] = 0
				stack = append(stack, temp[i])
				continue outer
			} else {
				cur := stack[len(stack)-1]
				if cur > temp[i] {
					result[i] = cur
					stack = append(stack, temp[i])
					continue outer
				} else {
					stack = stack[0 : len(stack)-1]
				}

			}

		}

	}
	return result

}
