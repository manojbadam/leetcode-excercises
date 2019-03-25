package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	capturePointers := make(map[*ListNode]int)
	for i := 0; headA != nil; i++ {
		capturePointers[headA] = i
		headA = headA.Next
	}
	for j := 0; headB != nil; j++ {
		if _, ok := capturePointers[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
