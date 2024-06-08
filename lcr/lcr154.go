package lcr

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	cache := map[*Node]*Node{}

	current := head

	for current != nil {
		cache[current] = &Node{
			Val: current.Val,
		}
		current = current.Next
	}

	current = head
	for current != nil {
		if current.Next != nil {
			cache[current].Next = cache[current.Next]
		}

		if current.Random != nil {
			cache[current].Random = cache[current.Random]
		}
		current = current.Next
	}
	return cache[head]

}
