package middle

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNodes(head *ListNode) *ListNode {
	cache := make([]*ListNode, 0)
	currentNode := head
outer:
	for currentNode != nil {
		for {
			if len(cache) == 0 {
				cache = append(cache, currentNode)
				currentNode = currentNode.Next
				continue outer
			}
			if cache[len(cache)-1].Val >= currentNode.Val {
				cache = append(cache, currentNode)
				currentNode = currentNode.Next
				continue outer
			} else {
				cache = cache[0 : len(cache)-1]
			}
		}
	}

	if len(cache)==1 {
		return cache[0]
	}

	for i:=1;i<len(cache);i++{
		cache[i-1].Next=cache[i]
		cache[i].Next=nil
	}
	return cache[0]

}
