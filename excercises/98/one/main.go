package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root != nil {
		if root.Left == nil && root.Right == nil {
			return true
		} else if root.Left == nil && root.Val < root.Right.Val {
			return isValidBST(root.Right)
		} else if root.Right == nil && root.Val > root.Left.Val {
			return isValidBST(root.Left)
		} else if root.Left != nil && root.Right != nil && root.Val > root.Left.Val && root.Val < root.Right.Val {
			return isValidBST(root.Left) && isValidBST(root.Right)
		} else {
			return false
		}
	} else {
		return true
	}
}
