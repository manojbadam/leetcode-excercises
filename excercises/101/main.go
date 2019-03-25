package solution

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Failing test case - [1,2]
func isSymmetric(root *TreeNode) bool {
	fmt.Println(root)
	if root == nil {
		return true
	} else if root.Left == nil && root.Right == nil {
		return true
	} else {
		return SymmetricHelper(root.Left, root.Right)
	}
}

func SymmetricHelper(nodeA *TreeNode, nodeB *TreeNode) bool {
	result := true
	if nodeA == nil && nodeB == nil {
		return true
	}
	if nodeA != nil && nodeB != nil && nodeA.Val == nodeB.Val {
		result = result && true
	} else if nodeA == nil && nodeB == nil {
		result = result && true
	} else {
		result = result && false
	}
	if nodeA.Left != nil && nodeB.Right != nil && nodeA.Left.Val == nodeB.Right.Val {
		result = result && true
	} else if nodeA.Left == nil && nodeB.Right == nil {
		result = result && true
	} else {
		result = result && false
	}
	fmt.Println("Outer", result)
	if nodeA.Right != nil && nodeB.Left != nil && nodeA.Right.Val == nodeB.Left.Val {
		result = result && true
	} else if nodeA.Right == nil && nodeB.Left == nil {
		result = result && true
	} else {
		result = result && false
	}
	return result && SymmetricHelper(nodeA.Left, nodeB.Right) && SymmetricHelper(nodeA.Right, nodeB.Left)
}
