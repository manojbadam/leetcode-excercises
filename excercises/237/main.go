package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	for true {
		node.Val = node.Next.Val
		if node.Next.Next == nil {
			node.Next = nil
			break
		} else {
			node = node.Next
		}
	}
}
