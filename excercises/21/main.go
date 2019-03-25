package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else {
		temp := &ListNode{Val: 0, Next: nil}
		head := temp
		for l1 != nil && l2 != nil {
			if l1.Val > l2.Val {
				temp.Next = l2
				temp = temp.Next
				l2 = l2.Next
			} else if l1.Val < l2.Val {
				temp.Next = l1
				temp = temp.Next
				l1 = l1.Next
			} else {
				temp.Next = l1
				temp = temp.Next
				l1 = l1.Next
				temp.Next = l2
				temp = temp.Next
				l2 = l2.Next
			}
		}
		if l1 != nil {
			temp.Next = l1
		} else {
			temp.Next = l2
		}
		return head.Next
	}
}
