package two

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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	tail := 0
	sum := 0
	var prevNode *ListNode
	var firstNode *ListNode
	for l1 != nil || l2 != nil {
		currentNode := ListNode{}
		sum = tail
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		currentNode.Val = sum % 10
		tail = sum / 10
		if prevNode != nil {
			prevNode.Next = &currentNode
		} else {
			firstNode = &currentNode
		}
		prevNode = &currentNode
	}
	if tail != 0 {
		tailNode := ListNode{}
		tailNode.Val = tail
		prevNode.Next = &tailNode
	}
	return firstNode
}
