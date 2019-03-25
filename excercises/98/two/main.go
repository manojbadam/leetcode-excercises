package solution

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	fmt.Println(root)
	if root == nil {
		return true
	}
	result := true
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left != nil {
		if root.Left.Val >= root.Val {
			return false
		} else {
			if root.Left.Left != nil {
				if root.Left.Left.Val >= root.Left.Val {
					return false
				}
			}
			if root.Left.Right != nil {
				if (root.Left.Right.Val <= root.Left.Val) || (root.Left.Right.Val >= root.Val) {
					return false
				}
			}
			result = true && isValidBST(root.Left)
		}
	}
	if root.Right != nil {
		if root.Right.Val <= root.Val {
			return false
		} else {
			if root.Right.Left != nil {
				if (root.Right.Left.Val >= root.Right.Val) || (root.Right.Left.Val <= root.Val) {
					return false
				}
			}
			if root.Right.Right != nil {
				if root.Right.Right.Val <= root.Right.Val {
					return false
				}
			}
			result = true && isValidBST(root.Right)
		}
	}
	return result
}
