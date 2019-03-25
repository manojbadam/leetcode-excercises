package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	} else {
		temp1 := head.Next
		temp2 := head.Next.Next
		head.Next = nil
		for head != nil && temp1 != nil {
			temp1.Next = head
			head = temp1
			temp1 = temp2
			if temp1 != nil {
				temp2 = temp1.Next
			}
		}
		return head
	}

}
