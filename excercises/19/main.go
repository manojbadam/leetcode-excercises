package nineteen

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n < 1 {
		return head
	} else {
		length := 0
		temp_head := head
		for temp_head != nil {
			temp_head = temp_head.Next
			length++
		}
		if length == 1 || length < n {
			return nil
		} else if length == n {
			head = head.Next
			return head
		} else {
			i := 0
			temp_head = head
			for i < length-n-1 {
				temp_head = temp_head.Next
				i++
			}
			temp_head.Next = temp_head.Next.Next
			return head
		}
	}
}
